package day04;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class ReposeRecord {
    public static void main(String[] args) {

    }

    private static List<String> getGuardRecords() throws IOException {
        return Files.readAllLines(Paths.get("src/day04/input.txt"));
    }


}
