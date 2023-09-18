package ru.uvuv643.helpers;

import ru.uvuv643.ontology.ItemData;

import java.util.List;
import java.util.stream.Collectors;

public class OntologyHelper {

    public static final String[] allowedItems = new String[]{
           "Helm_of_the_Overlord",
           "Helm_of_the_Dominator",
           "Vladmirs_Offering",
           "Helm_of_Iron_Will",
           "Diadem",
           "Ring_of_Basilius",
           "Sages_Mask",
           "Buckler",
           "Ring_of_Protection",
           "Morbid_Mask",
           "Blades_of_Attack",
           "Mask_of_Madness",
           "Quarterstaff",
           "Tango",
           "Bottle",
           "Helm_of_the_Overlord_Recipe",
           "Helm_of_the_Dominator_Recipe",
           "Ring_of_Basilius_Recipe",
           "Vladmirs_Offering_Recipe",
           "Buckler_Recipe",
    };

    public static ItemData getTheMostSimilarItem(String target) {
        List<String> theMostSimilarList = List.of(allowedItems)
            .stream()
            .filter((String item) -> item.toLowerCase().equals(target.toLowerCase()))
            .collect(Collectors.toList());
        if (theMostSimilarList.isEmpty()) {
            return null;
        }
        return new ItemData(theMostSimilarList.get(0));
    }

}
