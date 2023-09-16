package ru.uvuv643.parsers;

import ru.uvuv643.exceptions.ParseException;
import ru.uvuv643.exceptions.ValidateException;
import ru.uvuv643.io.RequestData;
import ru.uvuv643.ontology.ItemData;
import ru.uvuv643.validators.OntologyItemValidator;

import java.util.ArrayList;
import java.util.List;

public class DescriptionParser {

    private final List<ItemData> parsedDescription = new ArrayList<>();
    private final OntologyItemValidator validator = new OntologyItemValidator();

    public void parse(RequestData data) throws ParseException {

        String description = data.description;
        String[] elements = description.split(";");
        try {
            for (String item : elements) {
                ItemData itemData = new ItemData(item.trim().replace(' ', '_'));
                validator.validate(itemData);
                parsedDescription.add(itemData);
            }
        } catch (ValidateException exception) {
            throw new ParseException(exception.getMessage());
        }

    }

}
