package day03;

import java.awt.*;
import java.util.ArrayList;
import java.util.List;

public class FabricPiece {
    private Point _coordinates;
    private List<Integer> _claimants = new ArrayList<>();

    public FabricPiece(Point coordinates) {
        this._coordinates = coordinates;
    }

    public Point getCoordinates() {
        return _coordinates;
    }

    public List<Integer> getClaimants() {
        return _claimants;
    }

    public void addClaimant(int claimantId) {
        _claimants.add(claimantId);
    }
}
