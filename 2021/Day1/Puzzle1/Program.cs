var input = System.IO.File.ReadAllLines("input.txt");

int? previous = null;
var count = 0;
foreach (var line in input) {
    var current = int.Parse(line);
    if (previous != null) {
        if (current > previous) {
            count++;
        }
    }
    previous = current;
}

Console.WriteLine("Total Increments: " + count);