package ru.uvuv643.validators;

import ru.uvuv643.exceptions.ValidateException;
import ru.uvuv643.helpers.OntologyHelper;
import ru.uvuv643.ontology.ItemData;

import java.util.Set;

public class OntologyItemValidator {

    public void validate(String name) throws ValidateException {
        ItemData theMostSimilarInOntology = OntologyHelper.getTheMostSimilarItem(name);
        if (theMostSimilarInOntology == null) {
            throw new ValidateException("Указан предмет, которого нет в онтологии.");
        }
    }

}
