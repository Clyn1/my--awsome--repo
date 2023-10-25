#count regular files and directories in the current directories and subdirectories
count=$(find .  | wc -l)
#priny the count
echo $counts