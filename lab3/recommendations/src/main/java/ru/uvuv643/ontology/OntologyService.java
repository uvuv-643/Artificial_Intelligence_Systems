package ru.uvuv643.ontology;

import org.semanticweb.owlapi.apibinding.OWLManager;
import org.semanticweb.owlapi.model.*;
import org.semanticweb.HermiT.ReasonerFactory;
import org.semanticweb.owlapi.reasoner.Node;
import org.semanticweb.owlapi.reasoner.NodeSet;
import org.semanticweb.owlapi.reasoner.OWLReasoner;
import org.semanticweb.owlapi.reasoner.OWLReasonerFactory;
import org.semanticweb.owlapi.reasoner.structural.StructuralReasonerFactory;
import ru.uvuv643.helpers.OntologyHelper;

import java.io.File;
import java.util.*;
import java.util.stream.Collectors;

public class OntologyService {

    private final static String ONTOLOGY_PATH = "lab3/recommendations/ontology.rdf";
    private OWLOntology ontology;
    private OWLReasoner reasoner;


    public OntologyService() {
        try {
            OWLOntologyManager owlManager = OWLManager.createOWLOntologyManager();
            File file = new File(ONTOLOGY_PATH);
            this.ontology = owlManager.loadOntologyFromOntologyDocument(file);
        } catch (Exception exception) {
            System.err.println("Не найдено онтологии. Программа не будет дальше работать.");
            System.err.println("Загрузите онтологию ontology.rds в корневой каталог.");
            System.exit(-1);
        }
    }

    private OWLNamedIndividual getIndividualByName(String individualName) {
        for (OWLNamedIndividual individual : this.ontology.getIndividualsInSignature()) {
            if (individual.getIRI().getShortForm().equals(individualName)) {
                return individual;
            }
        }
        return null;
    }

    private Set<OWLNamedIndividual> getRelatedIndividuals(OWLNamedIndividual individual, String propertyName) {
        Set<OWLNamedIndividual> relatedIndividuals = new HashSet<>();
        OWLObjectProperty objectProperty = ontology.getOWLOntologyManager()
                .getOWLDataFactory()
                .getOWLObjectProperty(IRI.create(ontology.getOntologyID().getOntologyIRI().get() + "#" + propertyName));
        for (OWLNamedIndividual child : ontology.getIndividualsInSignature()) {
            if (reasoner.getObjectPropertyValues(individual, objectProperty).containsEntity(child)) {
                relatedIndividuals.add(child);
            }
        }
        return relatedIndividuals;
    }

    private Set<OWLClass> getTypes(OWLNamedIndividual individual) {
        Set<OWLClass> types = new HashSet<>();
        NodeSet<OWLClass> assertedTypes = reasoner.getTypes(individual, true);
        for (Node<OWLClass> assertedTypeNode : assertedTypes) {
            types.addAll(assertedTypeNode.getEntities());
        }
        return types;
    }

    public void giveRecommendations(List<ItemData> initial) {

        System.out.println("Вот список предметов которые вы ввели и их анализ");
        for (int i = 0; i < initial.size(); i++) {
            System.out.println((i + 1) + " -> " + initial.get(i).name.replace('_', ' '));
        }
        System.out.println();

        OWLReasonerFactory reasonerFactory = new ReasonerFactory();
        reasoner = reasonerFactory.createReasoner(ontology);
        reasoner.precomputeInferences();

        Set<String> canBeCreatedSet = new HashSet<>();

        for (ItemData individualData : initial) {

            System.out.println("Информация о предмете " + individualData.name.replace('_', ' '));
            String individualName = individualData.name;
            OWLNamedIndividual individual = this.getIndividualByName(individualName);
            Set<OWLNamedIndividual> parents = getRelatedIndividuals(individual, "parent");
            Set<OWLNamedIndividual> children = getRelatedIndividuals(individual, "child");
            Set<OWLClass> types = getTypes(individual);
            Set<String> typesInString = types.stream().map((OWLClass owlClass) -> owlClass.getIRI().getShortForm()).collect(Collectors.toSet());

            if (!parents.isEmpty()) {
                System.out.println("Вы можете создать такие предметы используя " + individualData.name.replace('_', ' ') + ":");
                for (OWLNamedIndividual parent : parents) {
                    System.out.println(parent.getIRI().getShortForm().replace('_', ' '));
                }
            } else {
                System.out.println("Вы ничего не можете собрать из этой штуки");
            }
            System.out.println();

            if (!children.isEmpty()) {
                System.out.println("Осмелюсь предположить, что вы собрали предмет " + individualData.name.replace('_', ' ') + " из следующих частей:");
                for (OWLNamedIndividual child : children) {
                    System.out.println(child.getIRI().getShortForm().replace('_', ' '));
                }
            } else {
                System.out.println("Вы не собирали этот предмет, а купили уже готовым в лавке");
            }
            System.out.println();

            if (typesInString.contains("CanBeDisassembled")) {
                System.out.println("Вы, кстати, можете разобрать этот предмет и вернуться к тем предметам, из которых его собрали!");
            } else {
                System.out.println("К сожалению, этот предмет разобрать не получится!");
            }
            System.out.println("-----------------------------");
            System.out.println();

            for (String targetIndividual : OntologyHelper.allowedItems) {
                Set<OWLNamedIndividual> targetChildren = getRelatedIndividuals(this.getIndividualByName(targetIndividual), "child");
                Set<String> childrenNames = targetChildren.stream().map((OWLNamedIndividual item) -> item.getIRI().getShortForm()).collect(Collectors.toSet());
                boolean canBeCreated = !targetChildren.isEmpty();
                if (!childrenNames.contains(individual.getIRI().getShortForm())) {
                    break;
                }
                if (canBeCreated) {
                    canBeCreatedSet.add(targetIndividual);
                }
            }

        }

        if (!canBeCreatedSet.isEmpty()) {
            System.out.println("Кроме того, из всех предметов, которые у вас есть, вы можете собрать: ");
            for (String item : canBeCreatedSet) {
                System.out.println(item.replace('_', ' '));
            }
        } else {
            System.out.println("Из тех предметов которые у вас есть вы ничего не можете собрать - не хватает вещей!");
        }

    }

}
