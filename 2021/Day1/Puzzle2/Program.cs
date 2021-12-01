using System.Collections.Generic;

var input = System.IO.File.ReadAllLines("input.txt");

var results = new List<int>();
for (int i = 0; i < input.Count()-2; i++) {
    var s1 = int.Parse(input[i]);
    var s2 = int.Parse(input[i+1]);
    var s3 = int.Parse(input[i+2]);
    results.Add(s1+s2+s3);
}

int? previous = null;
var count = 0;
foreach (var level in results) {
    if (previous != null) {
        if (level > previous) {
            count++;
        }
    }
    previous = level;
}

Console.WriteLine("Total Increments: " + count);
