// See https://aka.ms/new-console-template for more information

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

var distances = new List<int>();
for (var i = 0; i < list1.Count; i++)
{
    var dist = 0;
    if (list1[i] > list2[i])
    {
        dist = list1[i] - list2[i];
    }
    if (list2[i] > list1[i])
    {
        dist = list2[i] - list1[i];
    }
    distances.Add(dist);
}

var total = 0;
foreach (var dist in distances)
{
    total += dist;
}

Console.WriteLine(total);