#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <map>

using namespace std;

vector<string> readTravelDocs();
int findValidDocs(vector<string> travelDocs);

int main()
{
    vector<string> travelDocs = readTravelDocs();
    // for(int i = 0; i < travelDocs.size(); i++)
    // {
    //     cout << travelDocs[i] << endl;
    // }
    int validDocs = findValidDocs(travelDocs);
    cout << validDocs;

}

vector<string> readTravelDocs()
{
    ifstream fin;
    fin.open("input.txt");
    if (!fin.good())
    {
        cout << "File not opened";
        exit(1);
    }

    vector<string> travelDocs;
    while (!fin.eof())
    {
        string temp;
        getline(fin, temp);
        string doc = "";
        doc += temp;
        while (temp != "" && !fin.eof())
        {
            getline(fin, temp);
            if (temp != "")
            {
                doc += " " + temp;
            }
        }
        // cout << doc << endl;
        travelDocs.push_back(doc);
    }
    return travelDocs;
}

int findValidDocs(vector<string> travelDocs)
{
    int validDocs = 0;
    for (int i = 0; i < travelDocs.size(); i++)
    {
        map<string, string> docsMap;
        int index = 0;
        while (index != string::npos)
        {
            string key = travelDocs[i].substr(0,3);
            string value = travelDocs[i].substr(4,(travelDocs[i].find_first_of(" ") - 4));
            // cout << i << ": " << "key: " << key << " val: " << value << endl;
            index = travelDocs[i].find_first_of(" ");
            travelDocs[i] = travelDocs[i].substr(index + 1);
            docsMap[key] = value;
        }
        if(docsMap.size() == 8)
        {
            validDocs++;
        }
        else if(docsMap.size() == 7)
        {
            if(docsMap.count("cid") == 0)
            {
                validDocs++;
            }
        }
    }
    return validDocs;
}
