#include <iostream>
#include <string>
#include <fstream>
#include <vector>

using namespace std;

vector<string> makeGrid();
int findNumTrees(vector<string> grid);

int main()
{
    vector<string> grid;
    int numTrees;
    grid = makeGrid();
    numTrees = findNumTrees(grid);

    cout << numTrees;

}

vector<string> makeGrid()
{
    ifstream fin;
    vector<string> grid;

    fin.open("input.txt");
    if (!fin.good())
    {
        cout << "File not opened";
        exit(1);
    }

    while (!fin.eof())
    {
        string temp;
        fin >> temp;
        grid.push_back(temp);
    }

    return grid;
}

int findNumTrees(vector<string> grid)
{
    int numTrees = 0;
    int width = grid[0].size();
    int rightPos = 0;
    for(int d = 0; d < grid.size(); d++)
    {
        if(grid[d][rightPos] == '#' )
        {
            numTrees++;
        }
        rightPos = (rightPos + 3) % width;
    }
    return numTrees;
}