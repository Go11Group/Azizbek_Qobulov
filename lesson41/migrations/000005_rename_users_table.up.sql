-- Migration: Rename Memory to Storage in Phones Table
ALTER TABLE Phones
RENAME COLUMN Memory TO Storage;
