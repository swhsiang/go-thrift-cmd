namespace go computing

typedef i32 int

struct StatusOfService {
    1: required string version;
    // Listening which port
    2: required string network;
}

struct InputOfComputing {
    1: optional list<int> num_arr;
}

struct OutputOfComputing {
    1: required string error;
    2: optional int res;
}

service Computing {
    StatusOfService ping(),
    OutputOfComputing add(1:InputOfComputing input);
}
