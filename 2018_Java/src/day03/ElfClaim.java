package day03;

import java.awt.*;

public class ElfClaim {
    private int _claimId;
    private Point _startingCoordinates;
    private int _width;
    private int _height;

    public ElfClaim(String claimData) {
        String[] splitClaimData = claimData.split(" ");
        setClaimId(splitClaimData[0]);
        setStartingCoordinates(splitClaimData[2]);
        setWidthAndHeight(splitClaimData[3]);
    }

    public int getClaimId() {
        return _claimId;
    }

    public Point getStartingCoordinates() {
        return _startingCoordinates;
    }

    public int getWidth() {
        return _width;
    }

    public int getHeight() {
        return _height;
    }

    private void setClaimId(String unformattedData) {
        _claimId = Integer.parseInt(unformattedData.substring(1));
    }

    private void setStartingCoordinates(String unformattedData) {
        String[] splitCoordinates = unformattedData.split(",");
        _startingCoordinates = new Point(
                Integer.parseInt(splitCoordinates[0]),
                Integer.parseInt(splitCoordinates[1].substring(0, splitCoordinates[1].length() - 1))
        );
    }

    private void setWidthAndHeight(String unformattedData) {
        String[] splitWidthAndHeight = unformattedData.split("x");
        _width = Integer.parseInt(splitWidthAndHeight[0]);
        _height = Integer.parseInt(splitWidthAndHeight[1]);
    }
}
