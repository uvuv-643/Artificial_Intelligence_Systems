package ru.uvuv643.io;

import ru.uvuv643.io.interfaces.InputService;

import java.util.Scanner;

public class ConsoleInputService implements InputService {

    private static final Scanner scanner = new Scanner(System.in);

    public void greeting() {
        System.out.println("Привет!");
        System.out.println("Расскажи, какие предметы у тебя есть, а мы скажем, что ты можешь с ними сделать");
        System.out.println("Формат:");
        System.out.println("{предмет1}; {предмет2}; ...");
    }

    public RequestData input() {
        return new RequestData(scanner.next());
    }

}
