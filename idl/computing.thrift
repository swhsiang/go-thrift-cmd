namespace go swhsiang.computing

struct StatusOfService {
    1: required string version;
    // Listening which port
    2: required string network;
}

struct InputOfComputing {
    1: required list<i32> num_arr;
}

struct OutputOfComputing {
    1: required string error;
    2: optional i32 res;
}

service Computing {
    StatusOfService ping(),
    OutputOfComputing compute(1:InputOfComputing input);
}
