package com.penlook.service;

import redis.clients.jedis.Jedis;

public class Config {

	public Config() {
		Jedis redis = new Jedis("localhost");
	}
}
