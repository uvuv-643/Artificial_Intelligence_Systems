package ru.uvuv643;

import ru.uvuv643.exceptions.ParseException;
import ru.uvuv643.io.ConsoleInputService;
import ru.uvuv643.io.RequestData;
import ru.uvuv643.io.interfaces.InputService;
import ru.uvuv643.ontology.ItemData;
import ru.uvuv643.parsers.DescriptionParser;

import java.util.List;

public class Kernel {

    private static Kernel INSTANCE;

    private final DescriptionParser descriptionParser = new DescriptionParser();
    private final InputService inputService = new ConsoleInputService();


    private Kernel() {}

    public static Kernel getInstance() {
        if (INSTANCE == null) {
            INSTANCE = new Kernel();
        }
        return INSTANCE;
    }

    public void startInteractingViaConsole() {
        RequestData requestData = inputService.input();
        try {
            List<ItemData> itemDataList = descriptionParser.parse(requestData);
            System.out.println(itemDataList);
        } catch (ParseException exception) {
            try {
                System.err.println(exception.getMessage());
                System.err.println("Попробуйте ещё раз");
                System.err.println();
                Thread.sleep(1000);
                this.startInteractingViaConsole();
            } catch (InterruptedException ignored) { }
        }
    }

}
