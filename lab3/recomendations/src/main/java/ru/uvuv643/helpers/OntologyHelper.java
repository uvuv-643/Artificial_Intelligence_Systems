package ru.uvuv643.helpers;

import ru.uvuv643.ontology.ItemData;

import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class OntologyHelper {

    private static final ItemData[] allowedItems = new ItemData[]{
            new ItemData("Helm_of_the_Overlord"),
            new ItemData("Helm_of_the_Dominator"),
            new ItemData("Vladmirs_Offering"),
            new ItemData("Helm_of_Iron_Will"),
            new ItemData("Diadem"),
            new ItemData("Ring_of_Basilius"),
            new ItemData("Sages_Mask"),
            new ItemData("Buckler"),
            new ItemData("Ring_of_Protection"),
            new ItemData("Morbid_Mask"),
            new ItemData("Blades_of_Attack"),
            new ItemData("Mask_of_Madness"),
            new ItemData("Quarterstaff"),
            new ItemData("Tango"),
            new ItemData("Bottle"),
            new ItemData("Helm_of_the_Overlord_Recipe"),
            new ItemData("Helm_of_the_Dominator_Recipe"),
            new ItemData("Ring_of_Basilius_Recipe"),
            new ItemData("Vladmirs_Offering_Recipe"),
            new ItemData("Buckler_Recipe"),
    };

    public static Set<ItemData> getAllowedItems() {
        return new HashSet<>(List.of(allowedItems));
    }

}
