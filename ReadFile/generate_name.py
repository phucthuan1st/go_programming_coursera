from random import randint

fname = [
    "Jason",
    "John",
    "Kathy",
    "Mark",
    "Marry",
    "Taylor",
    "Jacob",
    "Nathaniel",
    "Michael",
]

lname = ["Greenwood", "Barry", "Kent", "Anderson", "Adam", "Messi", "Robert", "Ronald"]

with open("name.txt", "w") as f:
    n = len(fname)
    m = len(lname)

    lines = []

    for _ in range(n * m):
        line = f"{fname[randint(0, n - 1)]} {lname[randint(0, m - 1)]}\n"
        if line in lines:
            continue

        lines.append(line)
        f.write(line)
