#include <iostream>
#include <string>
#include <fstream>
#include <list>

using namespace std;

int main()
{
    ifstream fin;
    list<int> report;
    int expense;

    fin.open("input.txt");
    if (!fin.good())
    {
        cout << "File not opened";
        exit(1);
    }
    while (!fin.eof())
    {
        int temp;
        fin >> temp;
        report.push_back(temp);
    }
    for (list<int>::iterator it = report.begin(); it != report.end(); it++)
    {
        // cout << *it << endl;
        list<int>::iterator step = it;
        step++;

        for (step; step != report.end(); step++)
        {
            list<int>::iterator step2 = step;
            step2++;
            for (step2; step2 != report.end(); step2++)
            {
                if (*it + *step + *step2 == 2020)
                {
                    expense = *it * *step * *step2;
                    cout << *it << " " << *step << " " << *step2 << endl;
                }
            }
        }
    }
    cout << expense;
}
