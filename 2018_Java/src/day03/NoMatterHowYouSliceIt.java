package day03;

import java.awt.*;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.List;

public class NoMatterHowYouSliceIt {
    public static void main(String[] args) throws IOException {
        List<ElfClaim> elfClaims = getFormattedElfClaims();
        int maxGridWidth = getDimensionLength(elfClaims, "width");
        int maxGridHeight = getDimensionLength(elfClaims, "height");
        FabricPiece[][] fabricGrid = getFabricGrid(maxGridHeight, maxGridWidth);

        fabricGrid = setClaimants(fabricGrid, elfClaims);

        int overlapCount = getOverlapCount(fabricGrid);

        System.out.println(overlapCount);

        int intactClaimId = getIntactClaimId(fabricGrid, elfClaims);
        System.out.println(intactClaimId);
    }

    private static List<String> getElfClaims() throws IOException {
        return Files.readAllLines(Paths.get("src/day03/input.txt"));
    }

    private static List<ElfClaim> getFormattedElfClaims() throws IOException {
        var elfClaimDataSet = getElfClaims();
        List<ElfClaim> elfClaims = new ArrayList<>();
        for (String elfClaimData : elfClaimDataSet) {
            var elfClaim = new ElfClaim(elfClaimData);
            elfClaims.add(elfClaim);
        }

        return elfClaims;
    }

    private static int getDimensionLength(List<ElfClaim> elfClaims, String dimensionType) {
        int currentMax = 0;
        for (ElfClaim elfClaim : elfClaims) {
            int dimensionLength = dimensionType == "width" ?
                    elfClaim.getStartingCoordinates().x + elfClaim.getWidth() :
                    elfClaim.getStartingCoordinates().y + elfClaim.getHeight();

            if (dimensionLength > currentMax) {
                currentMax = dimensionLength;
            }
        }

        return currentMax;
    }

    private static FabricPiece[][] getFabricGrid(int height, int width) {
        var fabricGrid = new FabricPiece[height][width];

        for (int row = 0; row < fabricGrid.length; row++) {
            for (int column = 0; column < fabricGrid[0].length; column++) {
                var piece = new FabricPiece(new Point(column, row));
                fabricGrid[row][column] = piece;
            }
        }

        return fabricGrid;
    }

    private static FabricPiece[][] setClaimants(
            FabricPiece[][] fabricGrid,
            List<ElfClaim> elfClaims) {
        for (ElfClaim elfClaim : elfClaims) {
            Point cutInitialPosition = new Point(elfClaim.getStartingCoordinates().x, elfClaim.getStartingCoordinates().y);
            Point cutEndPosition = new Point(
                    elfClaim.getStartingCoordinates().x + elfClaim.getWidth() - 1,
                    elfClaim.getStartingCoordinates().y + elfClaim.getHeight() - 1);

            for (int row = cutInitialPosition.y; row <= cutEndPosition.y; row++) {
                for (int column = cutInitialPosition.x; column <= cutEndPosition.x; column++) {
                    fabricGrid[row][column].addClaimant(elfClaim.getClaimId());
                }
            }
        }

        return fabricGrid;
    }

    private static int getOverlapCount(FabricPiece[][] fabricGrid) {
        int overlapCount = 0;
        for (int row = 0; row < fabricGrid.length; row++) {
            for (int column = 0; column < fabricGrid[1].length; column++) {
                var piece = new FabricPiece(new Point(column, row));
                if (fabricGrid[row][column].getClaimants().size() > 1) {
                    overlapCount++;
                }
            }
        }

        return overlapCount;
    }

    private static int getIntactClaimId(
            FabricPiece[][] fabricGrid,
            List<ElfClaim> elfClaims) {
        for (ElfClaim elfClaim : elfClaims) {
            Point cutInitialPosition = new Point(elfClaim.getStartingCoordinates().x, elfClaim.getStartingCoordinates().y);
            Point cutEndPosition = new Point(
                    elfClaim.getStartingCoordinates().x + elfClaim.getWidth() - 1,
                    elfClaim.getStartingCoordinates().y + elfClaim.getHeight() - 1);

            boolean intactClaim = true;
            for (int row = cutInitialPosition.y; row <= cutEndPosition.y; row++) {
                for (int column = cutInitialPosition.x; column <= cutEndPosition.x; column++) {
                    if (fabricGrid[row][column].getClaimants().size() > 1) {
                        intactClaim = false;
                    }
                }
            }

            if (intactClaim) {
                return elfClaim.getClaimId();
            }
        }

        throw new IllegalArgumentException("An intact claimId exists! You just have to find it");
    }
}