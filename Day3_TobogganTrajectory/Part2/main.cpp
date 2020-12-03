#include <iostream>
#include <string>
#include <fstream>
#include <vector>

using namespace std;

vector<string> makeGrid();
int findNumTrees(vector<string> grid, int right, int down);

int main()
{
    vector<string> grid;
    int numTrees;
    int totTrees = 1;
    grid = makeGrid();
    char input;
    while(input != 'q'){
        int right;
        int down;
        cout << "Right" << endl;
        cout << "Down" << endl;
        cin >> down;
        numTrees = findNumTrees(grid, right, down);
        totTrees = totTrees * numTrees;
        cout << totTrees;
        cout << "Quit?" << endl;
        cin >> input;
    }

    cout << totTrees;

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

int findNumTrees(vector<string> grid, int right, int down)
{
    int numTrees = 0;
    int width = grid[0].size();
    int rightPos = 0;
    for(int d = 0; d < grid.size(); d += down)
    {
        if(grid[d][rightPos] == '#' )
        {
            numTrees++;
        }
        rightPos = (rightPos + right) % width;
    }
    return numTrees;
}