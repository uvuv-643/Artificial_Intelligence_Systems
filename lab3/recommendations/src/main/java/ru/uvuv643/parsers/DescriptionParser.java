package ru.uvuv643.parsers;

import ru.uvuv643.exceptions.ParseException;
import ru.uvuv643.exceptions.ValidateException;
import ru.uvuv643.helpers.OntologyHelper;
import ru.uvuv643.io.RequestData;
import ru.uvuv643.ontology.ItemData;
import ru.uvuv643.validators.OntologyItemValidator;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class DescriptionParser {

    private final OntologyItemValidator validator = new OntologyItemValidator();

    public List<ItemData> parse(RequestData data) throws ParseException {

        final List<ItemData> parsedDescription = new ArrayList<>();
        String description = data.description;
        String[] elements = description.split(";");
        try {
            for (String item : elements) {
                String name = item.trim().replace(' ', '_');
                validator.validate(name);
                parsedDescription.add(OntologyHelper.getTheMostSimilarItem(name));
            }
            return parsedDescription;
        } catch (ValidateException exception) {
            throw new ParseException(exception.getMessage());
        }

    }

}
