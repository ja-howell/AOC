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
    int low;
    int high;

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
        low = numPair.first;
        high = numPair.second;

        fin >> temp;
        let = temp[0];

        fin >> password;

        int letFound = 0;
        for (int i = 0; i < password.length(); i++)
        {

            if (password[i] == let)
            {
                letFound++;
            }
        }
        if (letFound >= low && letFound <= high)
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