#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <map>

using namespace std;

vector<string> readTravelDocs();
int findValidDocs(vector<string> travelDocs);
bool isValid(map<string, string> docMap);

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
            string key = travelDocs[i].substr(0, 3);
            string value = travelDocs[i].substr(4, (travelDocs[i].find_first_of(" ") - 4));
            // cout << i << ": " << "key: " << key << " val: " << value << endl;
            index = travelDocs[i].find_first_of(" ");
            travelDocs[i] = travelDocs[i].substr(index + 1);
            docsMap[key] = value;
        }
        if (docsMap.size() == 8)
        {
            if (isValid(docsMap))
            {
                validDocs++;
            }
        }
        else if (docsMap.size() == 7)
        {
            if (docsMap.count("cid") == 0)
            {
                if (isValid(docsMap))
                {
                    validDocs++;
                }
            }
        }
    }
    return validDocs;
}

bool isValid(map<string, string> docMap)
{
    if (stoi(docMap["byr"]) < 1920 || stoi(docMap["byr"]) > 2002)
    {
        // cout << stoi(docMap["byr"]) << endl;
        return false;
    }
    if (stoi(docMap["iyr"]) < 2010 || stoi(docMap["iyr"]) > 2020)
    {
        // cout << stoi(docMap["iyr"]) << endl;
        return false;
    }
    if (stoi(docMap["eyr"]) < 2020 || stoi(docMap["eyr"]) > 2030)
    {
        return false;
    }
    if (docMap["hgt"].find("cm") == string::npos && docMap["hgt"].find("in") == string::npos)
    {
        // cout << docMap["hgt"] << endl;
        return false;
    }
    else if (docMap["hgt"].find("cm") != string::npos)
    {
        int index = docMap["hgt"].find("cm");
        if (stoi(docMap["hgt"].substr(0, index)) < 150 || stoi(docMap["hgt"].substr(0, index)) > 193)
        {
            // cout << docMap["hgt"] << endl;
            return false;
        }
    }
    else if (docMap["hgt"].find("in") != string::npos)
    {
        int index = docMap["hgt"].find("in");
        if (stoi(docMap["hgt"].substr(0, index)) < 59 || stoi(docMap["hgt"].substr(0, index)) > 76)
        {
            // cout << docMap["hgt"] << endl;
            return false;
        }
    }
    if (docMap["hcl"][0] != '#')
    {
        // cout << docMap["hcl"][0] << endl;
        return false;
    }
    else if (docMap["hcl"].length() != 7)
    {
        // cout << docMap["hcl"] << endl;
        return false;
    }
    else
    {
        // cout << docMap["hcl"] << endl;
        for (int i = 1; i < docMap["hcl"].length(); i++)
        {
            if (!(docMap["hcl"][i] >= '0' && docMap["hcl"][i] <= '9') && !(docMap["hcl"][i] >= 'a' && docMap["hcl"][i] <= 'f'))
            {
                // cout << docMap["hcl"][i] << endl;
                return false;
            }
        }
    }
    string validEye[] = {"amb", "blu", "brn", "gry", "grn", "hzl", "oth"};
    bool eyeColValid = false;
    for (int i = 0; i < 7; i++)
    {
        if (docMap["ecl"] == validEye[i])
        {
            eyeColValid = true;
            break;
        }
    }
    if (!eyeColValid)
    {
        return false;
    }
    if (docMap["pid"].length() != 9)
    {
        return false;
    }
    else
    {
        for (int i = 1; i < docMap["pid"].length(); i++)
        {
            if (!(docMap["pid"][i] >= '0' && docMap["pid"][i] <= '9'))
            {
                return false;
            }
        }
    }
}