Head First SQL
by Lynn Beighley
Copyright © 2007 O’Reilly Media, Inc. All rights reserved.
Printed in the United States of America.
Published by O’Reilly Media, Inc., 1005 Gravenstein Highway North, Sebastopol, CA 95472.
O’Reilly Media books may be purchased for educational, business, or sales promotional use. Online editions are
also available for most titles (safari.oreilly.com). For more information, contact our corporate/institutional sales
department: (800) 998-9938 or corporate@oreilly.com.

Series Creators:

Kathy Sierra, Bert Bates

Series Editor:

Brett D. McLaughlin

Editor:

Catherine Nolan

Design Editor:

Louise Barr

Cover Designers:

Louise Barr, Karen Montgomery

Production Editor:

Sanders Kleinfeld

Indexer:

Julie Hawks

Page Viewer:

Andrew Fader

Printing History:
August 2007: First Edition.

He’s incredibly patient.
The O’Reilly logo is a registered trademark of O’Reilly Media, Inc. The Head First series designations,
Head First SQL, and related trade dress are trademarks of O’Reilly Media, Inc.
Many of the designations used by manufacturers and sellers to distinguish their products are claimed as
trademarks. Where those designations appear in this book, and O’Reilly Media, Inc., was aware of a trademark
claim, the designations have been printed in caps or initial caps.
While every precaution has been taken in the preparation of this book, the publisher and the authors assume no
responsibility for errors or omissions, or for damages resulting from the use of the information contained herein.
No clowns, doughnuts, or Girl Sprouts were harmed in the making of this book. Just my car, but it’s been fixed.
TM

This book uses RepKover™, a durable and ﬂexible lay-ﬂat binding.

ISBN-10: 0-596-52684-9
ISBN-13: 978-0-596-52684-9
[M]

table of contents

Table of Contents (Summary)
Intro

xxv

1

Data and Tables: A place for everything

1

2

The SELECT Statement: Gifted data retrieval

53

3

DELETE and UPDATE: A change will do you good

119

4

Smart Table Design: Why be normal?

159

5

ALTER: Rewriting the past

197

6

Advanced SELECT: Seeing your data with new eyes

235

7

Multi-table Database Design: Outgrowing your table

281

8

Joins and Multi-table Operations: Can’t we all just get along?

343

9

Subqueries: Queries Within Queries

379

10

Outer Joins, Self Joins, and Unions: New maneuvers

417

11

Constraints, Views, and Transactions: Too many cooks spoil the database

455

12

Security: Protecting your assets

493

Table of Contents (the real thing)
Intro
Your brain on SQL.  Here you are trying to learn something, while here
your brain is doing you a favor by making sure the learning doesn’t stick. Your
brain’s thinking, “Better leave room for more important things, like which wild
animals to avoid and whether naked snowboarding is a bad idea.” So how do you
trick your brain into thinking that your life depends on knowing SQL?

Who is this book for?

xxvi

We know what you’re thinking

xxvii

Metacognition

xxix

Bend your brain into submission

xxxi

Read me

xxxii

The technical review team

xxxiv

Acknowledgments

xxxv

ix

table of contents

1

data and tables
A place for everything
Don’t you just hate losing things? Whether it’s your car
keys, that 25% off coupon for Urban Outfitters, or your application’s
data, there’s nothing worse than not being able to keep up with what
you need... when you need it. And when it comes to your applications,
there’s no better place to store your important information than in a
table. So turn the page, come on in, and take a walk through the world
of relational databases.

Your database viewed
through x-ray specs...
Think of a database
like a container that
holds information…

column1
data

A table.

data
data
data

Another
table.

column2
data
data
data
data

data
data
data
data

column4

column5

data

column6

data

data

data

data

data

data

data

data

data

data

column1
data
data
data

column2
data

column3
data

column4

data

data

data

column1

column2

column3

data

data

data

data

data
data

column1
data
data
data

Another table.

7

What’s in a database?

8

Your database viewed through x-ray specs...

10

Databases contain connected data

12

Tables Up Close

13

Take command!

17

Setting the table: the CREATE TABLE statement

19

Creating a more complicated table

20

Look how easy it is to write SQL

21

Create the my_contacts table, finally

22

Your table is ready

23

Take a meeting with some data types

24

Your table, DESCribed

28

You can’t recreate an existing table or database!

30

Out with the old table, in with the new

32

To add data to your table, you’ll use the INSERT statement

34

Create the INSERT statement

37

Variations on an INSERT statement

41

Columns without values

42

Peek at your table with the SELECT statement

43

SQL Exposed: Confessions of a NULL

44

Controlling your inner NULL

45

NOT NULL appears in DESC

47

Fill in the blanks with DEFAULT

48

Your SQL Toolbox

50

data

data

data

2

Look at your data in categories

data

These are the columns.

These
are the

x

column3

Defining your data

column2
data
data

data

data
data
data

data
data
data

data

data

data

data

data

data

data

Some other table.

table of contents

2

the SELECT statement
Gifted data retrieval
Is it really better to give than retrieve? When it comes to
databases, chances are you’ll need to retrieve your data as often than
you’ll need to insert it. That’s where this chapter comes in: you’ll meet the
powerful SELECT statement and learn how to gain access to that important
information you’ve been putting in your tables. You’ll even learn how to
use WHERE, AND, and OR to selectively get to your data and even avoid
displaying the data that you don’t need.

I’m a star!

Date or no date?

54

A better SELECT

57

What the * is that?

58

How to query your data types

64

More punctuation problems

65

Unmatched single quotes

66

Single quotes are special characters

67

INSERT data with single quotes in it

68

SELECT specific columns to limit results

73

SELECT specific columns for faster results

73

Combining your queries

80

Finding numeric values

83

Smooth Comparison Operators

86

Finding numeric data with Comparison Operators

88

Text data roping with Comparison Operators

91

To be OR not to be

93

The difference between AND and OR

96

Use IS NULL to find NULLs

99

Saving time with a single keyword: LIKE

101

The call of the Wild(card)

101

Selecting ranges using AND and comparison operators

105

Just BETWEEN us… there’s a better way

106

After the dates, you are either IN...

109

... or you are NOT IN

110

More NOT

111

Your SQL Toolbox

116

xi

table of contents

3

DELETE and UPDATE
A change will do you good
Keep changing your mind? Now it’s OK! With the commands
you’re about to learn—DELETE and UPDATE—you’re no longer stuck with
a decision you made six months ago, when you first inserted that data about
mullets coming back into style soon. With UPDATE, you can change data, and
DELETE lets you get rid of data that you don’t need anymore. But we’re not just
giving you the tools; in this chapter, you’ll learn how to be selective with your new
powers and avoid dumping data that you really do need.

xii

Clowns are scary

120

Clown tracking

121

The clowns are on the move

122

How our clown data gets entered

126

Bonzo, we’ve got a problem

128

Getting rid of a record with DELETE

129

Using our new DELETE statement

131

DELETE rules

132

The INSERT-DELETE two step

135

Be careful with your DELETE

140

The trouble with imprecise DELETE

144

Change your data with UPDATE

146

UPDATE rules

147

UPDATE is the new INSERT-DELETE

148

UPDATE in action

149

Updating the clowns’ movements

152

UPDATE your prices

154

All we need is one UPDATE

156

Your SQL Toolbox

158

table of contents

4

smart table design
Why be normal?
You’ve been creating tables without giving much
thought to them. And that’s fine, they work. You can SELECT,

INSERT, DELETE, and UPDATE with them. But as you get more data,
you start seeing things you wish you’d done to make your WHERE

clauses simpler. What you need is to make your tables more normal.

Wait a second. I already have a table full of data.
You can't seriously expect me to use the DROP TABLE
command like I did in chapter 1 and type in all that data
again, just to create a primary key for each record…

Two fishy tables

160

A table is all about relationships

164

Atomic data

168

Atomic data and your tables

170

Atomic data rules

171

Reasons to be normal

174

The benefits of normal tables

175

Clowns aren’t normal

176

Halfway to 1NF

177

PRIMARY KEY rules

178

Getting to NORMAL

181

Fixing Greg’s table

182

The CREATE TABLE we wrote
table
Show me the money

183

Time-saving command

185

The CREATE TABLE with a PRIMARY KEY

186

1, 2, 3... auto incrementally

188

Adding a PRIMARY KEY to an existing table

192

ALTER TABLE and add a PRIMARY KEY

193

Your SQL Toolbox

194

184

xiii

table of contents

5

ALTER
Rewriting the Past
ver wished you could correct the mistakes of your past?
Well, now is your chance. By using the ALTER command, you can apply all the
lessons you’ve been learning to tables you designed days, months, even years ago.
Even better, you can do it without affecting your data. By the time you’re through
here, you’ll know what normal really means, and you’ll be able to apply it to all your
tables, past and present.

It’s time to turn your tired old
hooptie table into a date magnet
and take it to a level of table
pimpification you never knew existed.

xiv

We need to make some changes

198

Table altering

203

Extreme table makeover

204

Renaming the table

205

We need to make some plans

207

Retooling our columns

208

Structural changes

209

ALTER and CHANGE

210

Change two columns with one SQL statement

211

Quick! DROP that column

215

A closer look at the non-atomic location column

222

Look for patterns

223

A few handy string functions

224

Use a current column to fill a new column

229

How our UPDATE and SET combo works

230

Your SQL Toolbox

232

table of contents

6

advanced SELECT
Seeing your data with new eyes
It’s time to add a little finesse to your toolbox.  You already
know how to SELECT data and use WHERE clauses. But sometimes you need
more precision than SELECT and WHERE provide. In this chapter, you’ll learn
about how to order and group your data, as well as how to perform math
operations on your results.
Dataville Video is reorganizing

236

Problems with our current table

237

Matching up existing data

238

Populating the new column

239

UPDATE with a CASE expression

242

Looks like we have a problem

244

Tables can get messy

249

We need a way to organize the data we SELECT

250

Try a little ORDER BY

253

ORDER a single column

254

ORDER with two columns

257

ORDER with multiple columns

258

An orderly movietable

259

Reverse the ORDER with DESC

261

The Girl Sprout® cookie sales leader problem

263

SUM can add them for us

265

SUM all of them at once with GROUP BY

266

AVG with GROUP BY

267

MIN and MAX

268

COUNT the days

269

SELECT DISTINCT values

271

LIMIT the number of results

274

LIMIT to just second place

275

Your SQL Toolbox

278

xv

table of contents

7

multi-table database design
Outgrowing your table
Sometimes your single table isn’t big enough anymore.
 our data has become more complex, and that one table you’ve been using just
Y
isn’t cutting it. Your single table is full of redundant data, wasting space and
slowing down your queries. You’ve gone as far as you can go with a single table.
It’s a big world out there, and sometimes you need more than one table to
contain your data, control it, and ultimately, be the master of your own database.

interests

interests
int_id
interest

xvi

Finding Nigel a date

282

All is lost… But wait

293

Think outside of the single table

294

The multi-table clown tracking database

295

The clowntracking database schema

296

How to go from one table to two

298

Connecting your tables

303

Constraining your foreign key

305

Why bother with foreign keys?

306

CREATE a table with a FOREIGN KEY

307

Relationships between tables

309

Patterns of data: one‑to‑one

309

Patterns of data: when to use one‑to‑one tables

310

Patterns of data: one‑to‑many

311

Patterns of data: getting to many‑to‑many

312

Patterns of data: we need a junction table

315

Patterns of data: many-to-many

316

Finally in 1NF

321

Composite keys use multiple columns

322

Shorthand notations

324

Partial functional dependency

325

Transitive functional dependency

326

Second normal form

330

Third normal form (at last)

336

And so, Regis (and gregslist) lived happily ever after

339

Your SQL Toolbox

340

table of contents

8

joins and multi-table operations
Can’t we all just get along?
Welcome to a multi-table world. It’s great to have more than one table in
your database, but you’ll need to learn some new tools and techniques to work with
them. With multiple tables comes confusion, so you’ll need aliases to keep your tables
straight. And joins help you connect your tables, so that you can get at all the data you’ve
spread out. Get ready, it’s time to take control of your database again.

...and that’s where
little result tables
really come from.

Still repeating ourselves, still repeating...

344

Prepopulate your tables

345

We got the “table ain’t easy to normalize” blues

347

The special interests (column)

348

Keeping interested

349

UPDATE all your interests

350

Getting all the interests

351

Many paths to one place

352

CREATE, SELECT and INSERT at (nearly) the same time

352

CREATE, SELECT and INSERT at the same time

353

What’s up with that AS?

354

Column aliases

355

Table aliases, who needs ’em?

356

Everything you wanted to know about inner joins

357

Cartesian join

358

Releasing your inner join

363

The inner join in action: the equijoin

364

The inner join in action: the non-equijoin

367

The last inner join: the natural join

368

Joined-up queries?

375

Table and Column Aliases Exposed: What are you hiding from?

376

Your SQL Toolbox

377

xvii

table of contents

9

subqueries
Queries within queries
Yes, Jack, I’d like a two-part question, please. Joins are great,
but sometimes you need to ask your database more than one question. Or take
the result of one query and use it as the input to another query. That’s where
subqueries come in. They’ll help you avoid duplicate data, make your queries
more dynamic, and even get you in to all those high-end concert afterparties.
(Well, not really, but two out of three ain’t bad!)

Outer query

Greg gets into the job recruiting business

380

Greg’s list gets more tables

381

Greg uses an inner join

382

But he wants to try some other queries

384

Subqueries

386

We combine the two into a query with a subquery

387

As if one query wasn’t enough: meet the subquery

388

A subquery in action

389

Subquery rules

391

A subquery construction walkthrough

394

A subquery as a SELECT column

397

Another example: Subquery with a natural join

398

A noncorrelated subquery

399

SQL Exposed: Choosing the best way to query

400

A noncorrelated subquery with multiple values: IN, NOT IN

403

Correlated subqueries

408

A (useful) correlated subquery with NOT EXISTS

409

EXISTS and NOT EXISTS

410

Greg’s Recruiting Service is open for business

412

On the way to the party

413

Your SQL Toolbox

414

Inner query

xviii

table of contents

10

outer joins, self-joins, and unions
New maneuvers
You only know half of the story about joins. You’ve seen cross joins
that return every possible row, and inner joins that return rows from both tables where
there is a match. But what you haven’t seen are outer joins that give you back rows that
don’t have matching counterparts in the other table, self-joins which (strangely enough)
join a single table to itself, and unions that combine the results of queries. Once you
learn these tricks, you’ll be able to get at all your data exactly the way you need to. (And
we haven’t forgotten about exposing the truth about subqueries, either!)
Cleaning up old data

418

It’s about left and right

419

Here’s a left outer join

420

Outer joins and multiple matches

425

The right outer join

426

While you were outer joining…

429

We could create a new table

430

How the new table fits in

431

A self-referencing foreign key

432

Join the same table to itself

433

We need a self-join

435

Another way to get multi-table information

436

You can use a UNION

437

UNION is limited

438

UNION rules in action

439

UNION ALL

440

Create a table from your union

441

INTERSECT and EXCEPT

442

We’re done with joins, time to move on to…

443

Subqueries and joins compared

443

Turning a subquery into a join

444

A self-join as a subquery

449

Greg’s company is growing

450

Your SQL Toolbox

452

xix

table of contents

11

constraints, views, and transactions
Too many cooks spoil the database
Your database has grown and other people need to use it.
The problem is that some of them won’t be as skilled at SQL as you are. You need ways
to keep them from entering the wrong data, techniques for allowing them to only see
part of the data, and ways to stop them from stepping on each other when they try
entering data at the same time. In this chapter we begin protecting our data from the
mistakes of others. Welcome to Defensive Databases, Part 1.

Dataville
Savings & Loan

xx

Greg’s hired some help

456

Jim’s first day: Inserting a new client

457

Jim avoids a NULL

458

Flash forward three months

459

CHECK, please: Adding a CHECK CONSTRAINT

460

CHECKing the gender

461

Frank’s job gets tedious

463

Creating a view

465

Viewing your views

466

What your view is actually doing

467

What a view is

468

Inserting, updating, and deleting with views

471

The secret is to pretend a view is a real table

472

View with CHECK OPTION

475

Your view may be updatable if...

476

When you’re finished with your view

477

When bad things happen to good databases

478

What happened inside the ATM

479

More trouble at the ATM

480

It’s not a dream, it’s a transaction

482

The classic ACID test

483

SQL helps you manage your transactions

484

What should have happened inside the ATM

485

How to make transactions work with MySQL

486

Now try it yourself

487

Your SQL Toolbox

490

table of contents

12

security
Protecting your assets
You’ve put an enormous amount of time and energy into
creating your database. And you’d be devastated if anything happened to
it. You’ve also had to give other people access to your data, and you’re worried that
they might insert or update something incorrectly, or even worse, delete the wrong data.
You’re about to learn how databases and the objects in them can be made more secure,
and how you can have complete control over who can do what with your data.

root

bashful

doc

User problems

494

Avoiding errors in the clown tracking database

495

Protect the root user account

497

Add a new user

498

Decide exactly what the user needs

499

A simple GRANT statement

500

GRANT variations

503

REVOKE privileges

504

REVOKING a used GRANT OPTION

505

REVOKING with precision

506

The problem with shared accounts

510

Using your role

512

Role dropping

512

Using your role WITH ADMIN OPTION

514

Combining CREATE USER and GRANT

519

Greg’s List has gone global !

520

Your SQL Toolbox

522

How about a Greg’s List in your city?

524

Use SQL on your own projects and you too could be like Greg!

524

dopey

grumpy

happy

sleepy

sneezy

xxi

table of contents

i

leftovers
The Top Ten Topics (we didn’t cover)
Even after all that, there’s a bit more. There are just a few
more things we think you need to know. We wouldn’t feel right about ignoring
them, even though they only need a brief mention. So before you put the book
down, take a read through these short but important SQL tidbits.
Besides, once you’re done here, all that’s left is another appendix... and the
index... and maybe some ads... and then you’re really done. We promise!

A
B
C

D
E
F
G
H
I
J
K
L
M
N
O
P
Q
R
S
T
U
V
W

ABSOLUTE ACTION ADD ADMIN AFTER AGGREGATE ALIAS ALL ALLOCATE ALTER AND ANY ARE ARRAY AS
ASC ASSERTION AT AUTHORIZATION
BEFORE BEGIN BINARY BIT BLOB BOOLEAN BOTH BREADTH BY

CALL CASCADE CASCADED CASE CAST CATALOG CHAR CHARACTER CHECK CLASS CLOB CLOSE COLLATE
COLLATION COLUMN COMMIT COMPLETION CONNECT CONNECTION CONSTRAINT CONSTRAINTS
CONSTRUCTOR CONTINUE CORRESPONDING CREATE CROSS CUBE CURRENT CURRENT_DATE
CURRENT_PATH CURRENT_ROLE CURRENT_TIME CURRENT_TIMESTAMP CURRENT_USER CURSOR CYCLE

DATA DATE DAY DEALLOCATE DEC DECIMAL DECLARE DEFAULT DEFERRABLE DEFERRED DELETE DEPTH
DEREF DESC DESCRIBE DESCRIPTOR DESTROY DESTRUCTOR DETERMINISTIC DICTIONARY DIAGNOSTICS
DISCONNECT DISTINCT DOMAIN DOUBLE DROP DYNAMIC

EACH ELSE END END_EXEC EQUALS ESCAPE EVERY EXCEPT EXCEPTION EXEC EXECUTE EXTERNAL

FALSE FETCH FIRST FLOAT FOR FOREIGN FOUND FROM FREE FULL FUNCTION

GENERAL GET GLOBAL GO GOTO GRANT GROUP GROUPING
HAVING HOST HOUR

IDENTITY IGNORE IMMEDIATE IN INDICATOR INITIALIZE INITIALLY INNER INOUT INPUT INSERT
INT INTEGER INTERSECT INTERVAL INTO IS ISOLATION ITERATE
JOIN
KEY

LANGUAGE LARGE LAST LATERAL LEADING LEFT LESS LEVEL LIKE LIMIT LOCAL LOCALTIME
LOCALTIMESTAMP LOCATOR
MAP MATCH MINUTE MODIFIES MODIFY MODULE MONTH

NAMES NATIONAL NATURAL NCHAR NCLOB NEW NEXT NO NONE NOT NULL NUMERIC

OBJECT OF OFF OLD ON ONLY OPEN OPERATION OPTION OR ORDER ORDINALITY OUT OUTER OUTPUT

PAD PARAMETER PARAMETERS PARTIAL PATH POSTFIX PRECISION PREFIX PREORDER PREPARE
PRESERVE PRIMARY PRIOR PRIVILEGES PROCEDURE PUBLIC

READ READS REAL RECURSIVE REF REFERENCES REFERENCING RELATIVE RESTRICT RESULT RETURN
RETURNS REVOKE RIGHT ROLE ROLLBACK ROLLUP ROUTINE ROW ROWS

SAVEPOINT SCHEMA SCROLL SCOPE SEARCH SECOND SECTION SELECT SEQUENCE SESSION
SESSION_USER SET SETS SIZE SMALLINT SOME SPACE SPECIFIC SPECIFICTYPE SQL SQLEXCEPTION
SQLSTATE SQLWARNING START STATE STATEMENT STATIC STRUCTURE SYSTEM_USER
TABLE TEMPORARY TERMINATE THAN THEN TIME TIMESTAMP TIMEZONE_HOUR TIMEZONE_MINUTE TO
TRAILING TRANSACTION TRANSLATION TREAT TRIGGER TRUE

UNDER UNION UNIQUE UNKNOWN UNNEST UPDATE USAGE USER USING

VALUE VALUES VARCHAR VARIABLE VARYING VIEW

WHEN WHENEVER WHERE WITH WITHOUT WORK WRITE

X
Y
Z

YEAR

ZONE

xxii

#1. Get a GUI for your RDBMS

526

#2. Reserved Words and Special Characters

528

#3. ALL, ANY, and SOME

530

#4. More on Data Types

532

#5. Temporary tables

534

#6. Cast your data

535

#7. Who are you? What time is it?

536

#8. Useful numeric functions

537

#9. Indexing to speed things up

539

#10. 2-minute PHP/MySQL

540

table of contents

ii

iii

mySQL installation
Try it out for yourself
All your new SQL skills won’t do you much good
without a place to apply them. This appendix contains
instructions for getting your very own MySQL RDBMS for you to work with.
Get started, fast!

544

Instructions and Troubleshooting

544

Steps to Install MySQL on Windows

545

Steps to Install MySQL on Mac OS X

548

tools roundup
All your new SQL tools
Here are all your SQL tools in one place for the
first time, for one night only (kidding)! This is a
roundup of all the SQL tools we’ve covered. Take a moment to
survey the list and feel great—you learned them all!
Symbols

552

A–B

552

C–D

553

E–I

554

L–N

555

O–S

556

T–X

557

xxiii

