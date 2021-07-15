tokend tx token create-token  \
"abc" "ABC" "this is first token" "apple bus cat" 112000000 \
$(tokend keys show validator -a --home node1) true --home node1 --chain-id tokenchain-1