CREATE TABLE `kudos` (
	`KudoID`	TEXT NOT NULL UNIQUE,
	`KudoCount`	INTEGER DEFAULT 0,
	`URL`	TEXT,
	PRIMARY KEY(KudoID)
);
