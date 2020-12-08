#include <iostream>
#include <string>
#include <fstream>
#include <vector>

using namespace std;

vector<string> readPasses();
int findHighestSeatID(vector<string> passes);
int findSeatID(string pass);

int main()
{
    vector<string> passes = readPasses();
    int seatID = findHighestSeatID(passes);

    cout << seatID;
}

vector<string> readPasses()
{
    ifstream fin;
    vector<string> passes;

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
        passes.push_back(temp);
    }

    return passes;
}

int findHighestSeatID(vector<string> passes)
{
    int highestSeatID = 0;
    for (int i = 0; i < passes.size(); i++)
    {
        int seatID = findSeatID(passes[i]);

        if (seatID > highestSeatID)
        {
            highestSeatID = seatID;
        }
    }
    return highestSeatID;
}

int findSeatID(string pass)
{
    int colMax = 7;
    int colMin = 0;
    int rowMax = 127;
    int rowMin = 0;

    for (int j = 0; j < 10; j++)
    {
        if (pass[j] == 'F')
        {
            rowMax = (rowMax + rowMin) / 2;
        }
        else if (pass[j] == 'B')
        {
            rowMin = 1 + (rowMax + rowMin) / 2;
        }
        else if (pass[j] == 'R')
        {
            colMin = 1 + (colMax + colMin) / 2;
        }
        else
        {
            colMax = (colMax + colMin) / 2;
        }
    }

    return (rowMax * 8 + colMax);
}