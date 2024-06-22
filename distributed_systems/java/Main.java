import java.util.HashMap;
import java.util.Map;

public class Main {
    private static Map<String, String> dataStore = new HashMap<>();

    public static void main(String[] args) {
        addData("exampleKey", "exampleValue");
        System.out.println(getData("exampleKey"));
    }

    public static void addData(String key, String value) {
        dataStore.put(key, value);
    }

    public static String getData(String key) {
        return dataStore.getOrDefault(key, "Key not found!");
    }
}
