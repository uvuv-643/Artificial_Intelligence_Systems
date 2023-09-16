package ru.uvuv643;

public class Kernel {

    private static Kernel INSTANCE;

    private Kernel() {}

    public static Kernel getInstance() {
        if (INSTANCE == null) {
            INSTANCE = new Kernel();
        }
        return INSTANCE;
    }

    public void foo() {

    }

}
