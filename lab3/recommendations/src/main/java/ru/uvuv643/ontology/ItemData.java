package ru.uvuv643.ontology;

public class ItemData {

    public String name;

    public ItemData(String name) {
        this.name = name;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (this.getClass() == o.getClass()) {
            ItemData o1 = (ItemData) o;
            return this.name.equals(o1.name);
        }
        return false;
    }

    @Override
    public int hashCode() {
        return this.name.hashCode();
    }

}
