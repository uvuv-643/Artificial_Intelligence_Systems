package ru.uvuv643.io;

import ru.uvuv643.io.interfaces.InputService;

import java.util.Scanner;

public class ConsoleInputService implements InputService {

    private static final Scanner scanner = new Scanner(System.in);

    public void greeting() {
        System.out.println("Привет!");
        System.out.println("Расскажи, какие предметы у тебя есть, а мы скажем, что ты можешь с ними сделать. Формат вот такой:");
        System.out.println("{предмет1}; {предмет2}; ...");
        System.out.println("Печатай ниже:");
    }

    public RequestData input() {
        this.greeting();
        return new RequestData(scanner.nextLine().trim());
    }

}
