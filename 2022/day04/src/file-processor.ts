import * as fs from 'fs';
import AssignmentPair from './assignment-pair';

export default class FileProcessor {
  public static processInputIntoAssignmentPairs(
    inputPath: string
  ): AssignmentPair[] {
    return fs
      .readFileSync(inputPath)
      .toString()
      .split('\n')
      .map((rawAssignmentPair: string) =>
        this.createAssignmentPair(rawAssignmentPair)
      );
  }

  private static createAssignmentPair(
    rawAssignmentPair: string
  ): AssignmentPair {
    const assignmentA = {
      firstSection: parseInt(rawAssignmentPair.split(',')[0].split('-')[0]),
      lastSection: parseInt(rawAssignmentPair.split(',')[0].split('-')[1]),
    };
    const assignmentB = {
      firstSection: parseInt(rawAssignmentPair.split(',')[1].split('-')[0]),
      lastSection: parseInt(rawAssignmentPair.split(',')[1].split('-')[1]),
    };

    return new AssignmentPair(assignmentA, assignmentB);
  }
}
