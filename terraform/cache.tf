resource "aws_elasticache_serverless_cache" "prompt-defender-cache" {
  engine = "redis"
  name   = "prompt-defender-cache"

  cache_usage_limits {
    data_storage {
      maximum = 10
      unit    = "GB"
    }
    ecpu_per_second {
      maximum = 5000
    }
  }
  description          = "Cache used for Prompt Defender"
  major_engine_version = "7"
}