package day01;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

public class ChronalCalibration {

    public static void main(String[] args) throws IOException {
        List<String> input = getInput();

        List<Integer> usedFrequencies = new ArrayList<>();
        int currentFrequency = 0;
        int currentInputIndex = 0;

        while (!usedFrequencies.contains(currentFrequency)) {
            usedFrequencies.add(currentFrequency);
            currentFrequency += Integer.parseInt(input.get(currentInputIndex % (input.size())));
            currentInputIndex++;
        }

        System.out.println(currentFrequency);
    }

    private static List<String> getInput() throws IOException {
        return Files.readAllLines(Paths.get("src/day01/input.txt"));
    }
}
