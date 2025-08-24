awk '
$1 ~ /^BenchmarkDay[0-9]+Part[0-9]+-/ {
  name=$1; sub(/-.*/, "", name)
  for (i=1;i<=NF;i++) if ($i=="ns/op") { val=$(i-1)+0 }
  if (!(name in min) || val < min[name]) min[name]=val
}
END { total=0; for (k in min) total+=min[k]; printf("Min-sum: %.0f ns (%.3f ms, %.6f s)\n", total, total/1e6, total/1e9) }'
