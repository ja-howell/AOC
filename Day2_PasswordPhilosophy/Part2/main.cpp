#include <iostream>
#include <string>
#include <fstream>
#include <utility>

using namespace std;
pair<int, int> ParseBounds(string numLets);

int main()
{
    ifstream fin;
    string numLets;
    char let;
    string password;
    int pos1;
    int pos2;

    int numGoodPasswords = 0;

    fin.open("input.txt");
    if (!fin.good())
    {
        cout << "File not opened";
        exit(1);
    }
    while (!fin.eof())
    {
        string temp;

        fin >> numLets;
        pair<int, int> numPair = ParseBounds(numLets);
        pos1 = numPair.first;
        pos2 = numPair.second;

        fin >> temp;
        let = temp[0];

        fin >> password;

        // ^ is xor operator
        if((password[pos1-1] == let) ^ (password[pos2-1] == let))
        {
            numGoodPasswords++;
        }

    }
    cout << numGoodPasswords;
}

pair<int, int> ParseBounds(string numLets)
{
    int dashIndex = numLets.find("-");
    // cout << dashIndex << endl;

    string temp1 = numLets.substr(0, dashIndex);
    string temp2 = numLets.substr(dashIndex + 1);

    return make_pair(stoi(temp1), stoi(temp2));
}