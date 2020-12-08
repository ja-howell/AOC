#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <algorithm>

using namespace std;

vector<string> readPasses();
int findSeatID(string pass);

int main()
{
    vector<string> passes = readPasses();
    vector<int> seatIDs;
    int mySeatID;
    for (int i = 0; i < passes.size(); i++)
    {
        seatIDs.push_back(findSeatID(passes[i]));
    }

    sort(seatIDs.begin(), seatIDs.end());

    for(int i = 0; i < seatIDs.size() - 1; i++)
    {
        int preSeat = seatIDs[i];
        int postSeat = seatIDs[i+1];

        if(postSeat != preSeat + 1)
        {
            mySeatID = preSeat + 1;
            break;
        }        
    } 


    cout << mySeatID;
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