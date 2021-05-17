package day02;

import java.io.IOException;
import java.lang.reflect.Array;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;

public class InventoryManagementSystem {

    private static final char[] alphabet = {
            'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
            'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'
    };

    public static void main(String[] args) throws IOException {
        List<String> boxIds = getBoxIds();
        int twoLetterCount = 0;
        int threeLetterCount = 0;

        for (int i = 0; i < boxIds.size(); i++) {
            if (hasOccurrenceCount(boxIds.get(i), 2)) {
                twoLetterCount++;
            }

            if (hasOccurrenceCount(boxIds.get(i), 3)) {
                threeLetterCount++;
            }
        }

        System.out.println(twoLetterCount * threeLetterCount);

        String[] correctBoxIds = getCorrectBoxIds(boxIds);
        String commonLetters = getCommonLetters(correctBoxIds);

        System.out.println(commonLetters);
    }

    private static List<String> getBoxIds() throws IOException {
        return Files.readAllLines(Paths.get("src/day02/input.txt"));
    }

    private static boolean hasOccurrenceCount(String input, int countTarget) {
        for (int i = 0; i < alphabet.length; i++) {

            int letterOccurrenceCount = 0;
            for (int j = 0; j < input.length(); j++) {
                if (input.charAt(j) == alphabet[i]) {
                    letterOccurrenceCount++;
                }
            }
            
            if (letterOccurrenceCount == countTarget) {
                return true;
            }
        }

        return false;
    }

    private static String getCommonLetters(String[] boxIds) {
        var commonLetters = new StringBuilder();
        for (int i = 0; i < boxIds[0].length(); i++) {
            if (boxIds[0].charAt(i) == boxIds[1].charAt(i)) {
                commonLetters.append(boxIds[0].charAt(i));
            }
        }

        return commonLetters.toString();
    }

    private static String[] getCorrectBoxIds(List<String> boxIds) {
        String[] correctBoxIds = new String[2];
        for (int i = 0; i < boxIds.size(); i++) {
            for (int j = 0; j < boxIds.size(); j++) {
                if (differsByOneCharacter(boxIds.get(i), boxIds.get(j))) {
                    correctBoxIds[0] = boxIds.get(i);
                    correctBoxIds[1] = boxIds.get(j);

                    return correctBoxIds;
                }
            }
        }

        return correctBoxIds;
    }

    private static boolean differsByOneCharacter(String boxId1, String boxId2) {
        int differenceCount = 0;
        for (int i = 0; i < boxId1.length(); i++) {
            if (boxId1.charAt(i) != boxId2.charAt(i)) {
                differenceCount++;
            }
        }

        return differenceCount == 1;
    }
}
