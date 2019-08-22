public class Foo {
    public static void main(String[] args) {
        System.out.println("Hello, world.");    

        Person p = new Person();
        System.out.println(p.isNull()); // => false

        Person p2 = null;
        System.out.println(p2.isNull());
        // Exception in thread "main" java.lang.NullPointerException
        // 	at Foo.main(Foo.java:9)
    } 
}

class Person {
    boolean isNull() {
        return this == null;
    }
}
