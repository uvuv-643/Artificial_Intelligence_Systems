package ru.uvuv643.validators;

import ru.uvuv643.exceptions.ValidateException;
import ru.uvuv643.helpers.OntologyHelper;
import ru.uvuv643.ontology.ItemData;

import java.util.Set;

public class OntologyItemValidator {

    private static final Set<ItemData> allowedItems = OntologyHelper.getAllowedItems();

    public void validate(ItemData itemData) throws ValidateException {
        if (!allowedItems.contains(itemData)) {
            throw new ValidateException("Указан предмет, которого нет в онтологии.");
        }
    }

}
