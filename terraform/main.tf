resource "scaleway_function" "main" {
  name         = "overcelebrate-fn"
  namespace_id = scaleway_function_namespace.ns.id
  runtime      = "go118"
  handler      = "StartFn"
  privacy      = "public"
  zip_file     = "func3.zip"
  min_scale = 1
  max_scale = 1
  memory_limit = 128
  description = "the overcelebrate fn"
}

resource "scaleway_function_namespace" "ns" {
  name        = "overcelebrate-ns"
  description = "overcelebrate-ns"
}
