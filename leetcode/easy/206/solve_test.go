package _206

import (
    "github.com/zhangli2946/tiantai/common"
    "testing"
)

func TestSolve(t *testing.T) {
    s := reverseList2(&common.ListNode{
        1,
        &common.ListNode{
            2,
            &common.ListNode{
                3,
                &common.ListNode{
                    4,
                    &common.ListNode{
                        5, nil,
                    },
                },
            },
        },
    })
    t.Log(s)

}
