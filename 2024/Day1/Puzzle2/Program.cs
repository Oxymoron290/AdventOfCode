var list1 = new List<int>();
var list2 = new List<int>();
var input = File.ReadAllLines("input.txt").ToList();
foreach (var line in input)
{
    var lineParts = line.Split("  ");
    list1.Add(int.Parse(lineParts[0]));
    list2.Add(int.Parse(lineParts[1]));
}

list1.Sort();
list2.Sort();

var similarities = new List<int>();
foreach(var item in list1)
{
    var count = list2.Count(x => x == item);
    Console.WriteLine($"x: {count}");
    similarities.Add(item * count);
}

var total = 0;
foreach (var s in similarities)
{
    total += s;
}

Console.WriteLine(total);