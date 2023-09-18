package ru.uvuv643;

import org.semanticweb.owlapi.apibinding.OWLManager;
import org.semanticweb.owlapi.model.IRI;
import org.semanticweb.owlapi.model.OWLOntology;
import org.semanticweb.owlapi.model.OWLOntologyCreationException;
import org.semanticweb.owlapi.model.OWLOntologyManager;

import java.io.File;

public class Main {

    public Kernel kernel = Kernel.getInstance();

    public static void main(String[] args) throws OWLOntologyCreationException {

        Kernel.getInstance().startInteractingViaConsole();

    }
}