import AssignmentPair from './assignment-pair';

export default class TestData {
  public static assignmentPairs = [
    new AssignmentPair(
      { firstSection: 2, lastSection: 4 },
      { firstSection: 6, lastSection: 8 }
    ),
    new AssignmentPair(
      { firstSection: 2, lastSection: 3 },
      { firstSection: 5, lastSection: 7 }
    ),
    new AssignmentPair(
      { firstSection: 5, lastSection: 7 },
      { firstSection: 7, lastSection: 9 }
    ),
    new AssignmentPair(
      { firstSection: 2, lastSection: 8 },
      { firstSection: 3, lastSection: 7 }
    ),
    new AssignmentPair(
      { firstSection: 6, lastSection: 6 },
      { firstSection: 4, lastSection: 6 }
    ),
    new AssignmentPair(
      { firstSection: 2, lastSection: 6 },
      { firstSection: 4, lastSection: 8 }
    ),
  ];
}
