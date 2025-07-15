package main

import "time"

type Data struct {
    Timestamp time.Time
}

func (r *RedisRepo) Write(data Data, key string) error {
    conn := r.pool.Get()
    defer conn.Close()
    _, err := conn.Do("HSET", key, "timestamp", data.Timestamp.Unix())
    return err
}

func (r *RedisRepo) Read(key string) (*Data, error) {
    conn := r.pool.Get()
    defer conn.Close()
    v, err := redis.Values(conn.Do("HGETALL", key))
    if err != nil {
        return nil, err
    }

    var fields struct {
        Timestamp int64 `redis:"timestamp"`
    }

    err = redis.ScanStruct(v, &fields)
    if err != nil {
        return nil, err
    }
    return &Data{Timestamp: time.Unix(fields.Timestamp, 0)}, nil
}

///

type Data struct {
    Timestamp time.Time
}

func (r *RedisRepo) Write(data Data, key string) error {
    conn := r.pool.Get()
    defer conn.Close()
    _, err := conn.Do("HSET", key, "timestamp", data.Timestamp.Format(time.RFC3339))
    return err
}

func (r *RedisRepo) Read(key string) (*Data, error) {
    conn := r.pool.Get()
    defer conn.Close()
    v, err := redis.Values(conn.Do("HGETALL", key))
    if err != nil {
        return nil, err
    }

    var fields struct {
        Timestamp string `redis:"timestamp"`
    }

    err = redis.ScanStruct(v, &fields)
    if err != nil {
        return nil, err
    }
    t, err := time.Parse(time.RFC3339, fields.Timestamp)
    if err != nil {
        return nil, err
    }
    return &Data{Timestamp: t}, nil
}
