package ru.uvuv643.exceptions;

public class ValidateException extends Exception {

    private String message;

    public ValidateException(String message) {
        this.message = message;
    }

    public String getMessage() {
        return this.message;
    }

}
