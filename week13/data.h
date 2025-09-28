#include<iostream>

using namespace std;

class circle;
class box {
	int color;
	int upx, upy, lx, ly;
public:
	friend int same_color(circle c, box b);
	void set_color(int color) {
		this->color = color;
	}
	int get_color() {re}
};
class circle {
	int color;
	int x, y, r;
public:
	void set_color(int color) {
		this->color = color;
	}
	void set_x(int x) { this->x = x; }
	void set_y(int a) { y=a; }
	void set_r(int r) { this->r = r; }
	int get_color() { return color; }
	int get_x() { return x; }
	int get_y() { return y; }
	int get_r() { return r; }
};

int same_color(circle c, box b);