#include "compare_length.h"

// #include <iostream>

int compare(const std::vector<int> &vl, const std::vector<int> &vr) {
  int l2_l = l2(vl);
  int l2_r = l2(vr);
  return l2_l - l2_r;
}

// int main(int argc, char *argv[]) {
//   std::vector<int> v1(5, 0);
//   std::vector<int> v2(4, 1);
//   std::cout << "compare value is " << compare(v1, v2) << std::endl;
// }