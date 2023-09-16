package ru.uvuv643.exceptions;

public class ParseException extends Exception {

    private String message;

    public ParseException(String message) {
        this.message = message;
    }

    public String getMessage() {
        return this.message;
    }

}
