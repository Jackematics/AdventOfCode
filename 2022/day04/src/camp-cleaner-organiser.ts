import AssignmentPair from './assignment-pair';

export default class CampCleanupOrganiser {
  public static calculateTotalFullSectionDuplication(
    assignmentPairs: AssignmentPair[]
  ): number {
    return assignmentPairs.reduce(
      (acc: number, assignmentPair: AssignmentPair) =>
        acc + assignmentPair.incrementIfSectionFullyDuplicated(),
      0
    );
  }

  public static calculateTotalPartialSectionDuplication(
    assignmentPairs: AssignmentPair[]
  ): number {
    return assignmentPairs.reduce(
      (acc: number, assignmentPair: AssignmentPair) =>
        acc + assignmentPair.incrementIfSectionPartiallyDuplicated(),
      0
    );
  }
}
