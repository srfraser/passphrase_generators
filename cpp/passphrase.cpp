#include <fstream>
#include <vector>
#include <string>
#include <cstdlib>
#include <ctime>

using namespace std;

int main(int argc, char* argv[])
{
        auto length = 4;
        ifstream f ("/usr/share/dict/words");
        if (!f.is_open())
                perror("error while opening file");

        string line;
        std::vector<string> wordlist;
        while(getline(f, line)) {
                wordlist.push_back(line);
        }

        if (f.bad())
                perror("error while reading file");

        srand(time(nullptr));
        for (;length > 0; length--) {
            printf("%s ", wordlist[rand() % wordlist.size()].c_str());
        }
        printf("\n");
        return 0;
}
