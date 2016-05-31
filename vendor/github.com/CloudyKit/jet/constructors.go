// Copyright 2016 José Santos <henrique_1609@me.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package jet

import (
	"fmt"
	"strconv"
	"strings"
)

func (t *Template) newSliceExpr(pos Pos, line int, base, index, len Expression) *SliceExprNode {
	return &SliceExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeSliceExpr, Pos: pos, Line: line}, Index: index, Base: base, EndIndex: len}
}

func (t *Template) newIndexExpr(pos Pos, line int, base, index Expression) *IndexExprNode {
	return &IndexExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeIndexExpr, Pos: pos, Line: line}, Index: index, Base: base}
}

func (t *Template) newTernaryExpr(pos Pos, line int, boolean, left, right Expression) *TernaryExprNode {
	return &TernaryExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeTernaryExpr, Pos: pos, Line: line}, Boolean: boolean, Left: left, Right: right}
}

func (t *Template) newBuiltinExpr(pos Pos, line int, name string, nodetype NodeType) *BuiltinExprNode {
	return &BuiltinExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: nodetype, Pos: pos, Line: line}, Name: name}
}

func (t *Template) newSet(pos Pos, line int, isLet bool, left, right []Expression) *SetNode {
	return &SetNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeSet, Pos: pos, Line: line}, Let: isLet, Left: left, Right: right}
}

func (t *Template) newCallExpr(pos Pos, line int, expr Expression) *CallExprNode {
	return &CallExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeCallExpr, Pos: pos, Line: line}, BaseExpr: expr}
}
func (t *Template) newNotExpr(pos Pos, line int, expr Expression) *NotExprNode {
	return &NotExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeNotExpr, Pos: pos, Line: line}, Expr: expr}
}
func (t *Template) newNumericComparativeExpr(pos Pos, line int, left, right Expression, item item) *NumericComparativeExprNode {
	return &NumericComparativeExprNode{binaryExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeNumericComparativeExpr, Pos: pos, Line: line}, Operator: item, Left: left, Right: right}}
}

func (t *Template) newComparativeExpr(pos Pos, line int, left, right Expression, item item) *ComparativeExprNode {
	return &ComparativeExprNode{binaryExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeComparativeExpr, Pos: pos, Line: line}, Operator: item, Left: left, Right: right}}
}

func (t *Template) newLogicalExpr(pos Pos, line int, left, right Expression, item item) *LogicalExprNode {
	return &LogicalExprNode{binaryExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeLogicalExpr, Pos: pos, Line: line}, Operator: item, Left: left, Right: right}}
}

func (t *Template) newMultiplicativeExpr(pos Pos, line int, left, right Expression, item item) *MultiplicativeExprNode {
	return &MultiplicativeExprNode{binaryExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeMultiplicativeExpr, Pos: pos, Line: line}, Operator: item, Left: left, Right: right}}
}

func (t *Template) newAdditiveExpr(pos Pos, line int, left, right Expression, item item) *AdditiveExprNode {
	return &AdditiveExprNode{binaryExprNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeAdditiveExpr, Pos: pos, Line: line}, Operator: item, Left: left, Right: right}}
}

func (t *Template) newList(pos Pos) *ListNode {
	return &ListNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeList, Pos: pos}}
}

func (t *Template) newText(pos Pos, text string) *TextNode {
	return &TextNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeText, Pos: pos}, Text: []byte(text)}
}

func (t *Template) newPipeline(pos Pos, line int) *PipeNode {
	return &PipeNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodePipe, Pos: pos, Line: line}}
}

func (t *Template) newAction(pos Pos, line int) *ActionNode {
	return &ActionNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeAction, Pos: pos, Line: line}}
}

func (t *Template) newCommand(pos Pos) *CommandNode {
	return &CommandNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeCommand, Pos: pos}}
}

func (t *Template) newNil(pos Pos) *NilNode {
	return &NilNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeNil, Pos: pos}}
}

func (t *Template) newField(pos Pos, ident string) *FieldNode {
	return &FieldNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeField, Pos: pos}, Ident: strings.Split(ident[1:], ".")} // [1:] to drop leading period
}

func (t *Template) newChain(pos Pos, node Node) *ChainNode {
	return &ChainNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeChain, Pos: pos}, Node: node}
}

func (t *Template) newBool(pos Pos, true bool) *BoolNode {
	return &BoolNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeBool, Pos: pos}, True: true}
}

func (t *Template) newString(pos Pos, orig, text string) *StringNode {
	return &StringNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeString, Pos: pos}, Quoted: orig, Text: text}
}

func (t *Template) newEnd(pos Pos) *endNode {
	return &endNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: nodeEnd, Pos: pos}}
}

func (t *Template) newElse(pos Pos, line int) *elseNode {
	return &elseNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: nodeElse, Pos: pos, Line: line}}
}

func (t *Template) newIf(pos Pos, line int, set *SetNode, pipe Expression, list, elseList *ListNode) *IfNode {
	return &IfNode{BranchNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeIf, Pos: pos, Line: line}, Set: set, Expression: pipe, List: list, ElseList: elseList}}
}

func (t *Template) newRange(pos Pos, line int, set *SetNode, pipe Expression, list, elseList *ListNode) *RangeNode {
	return &RangeNode{BranchNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeRange, Pos: pos, Line: line}, Set: set, Expression: pipe, List: list, ElseList: elseList}}
}

func (t *Template) newBlock(pos Pos, line int, name string, pipe Expression, listNode *ListNode) *BlockNode {
	return &BlockNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeBlock, Line: line, Pos: pos}, Name: name, Expression: pipe, List: listNode}
}

func (t *Template) newYield(pos Pos, line int, name string, pipe Expression) *YieldNode {
	return &YieldNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeYield, Pos: pos, Line: line}, Name: name, Expression: pipe}
}

func (t *Template) newInclude(pos Pos, line int, name string, pipe Expression) *IncludeNode {
	return &IncludeNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeInclude, Pos: pos, Line: line}, Name: name, Expression: pipe}
}

func (t *Template) newNumber(pos Pos, text string, typ itemType) (*NumberNode, error) {
	n := &NumberNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeNumber, Pos: pos}, Text: text}
	switch typ {
	case itemCharConstant:
		rune, _, tail, err := strconv.UnquoteChar(text[1:], text[0])
		if err != nil {
			return nil, err
		}
		if tail != "'" {
			return nil, fmt.Errorf("malformed character constant: %s", text)
		}
		n.Int64 = int64(rune)
		n.IsInt = true
		n.Uint64 = uint64(rune)
		n.IsUint = true
		n.Float64 = float64(rune) // odd but those are the rules.
		n.IsFloat = true
		return n, nil
	case itemComplex:
		// fmt.Sscan can parse the pair, so let it do the work.
		if _, err := fmt.Sscan(text, &n.Complex128); err != nil {
			return nil, err
		}
		n.IsComplex = true
		n.simplifyComplex()
		return n, nil
	}
	// Imaginary constants can only be complex unless they are zero.
	if len(text) > 0 && text[len(text)-1] == 'i' {
		f, err := strconv.ParseFloat(text[:len(text)-1], 64)
		if err == nil {
			n.IsComplex = true
			n.Complex128 = complex(0, f)
			n.simplifyComplex()
			return n, nil
		}
	}
	// Do integer test first so we get 0x123 etc.
	u, err := strconv.ParseUint(text, 0, 64) // will fail for -0; fixed below.
	if err == nil {
		n.IsUint = true
		n.Uint64 = u
	}
	i, err := strconv.ParseInt(text, 0, 64)
	if err == nil {
		n.IsInt = true
		n.Int64 = i
		if i == 0 {
			n.IsUint = true // in case of -0.
			n.Uint64 = u
		}
	}
	// If an integer extraction succeeded, promote the float.
	if n.IsInt {
		n.IsFloat = true
		n.Float64 = float64(n.Int64)
	} else if n.IsUint {
		n.IsFloat = true
		n.Float64 = float64(n.Uint64)
	} else {
		f, err := strconv.ParseFloat(text, 64)
		if err == nil {
			// If we parsed it as a float but it looks like an integer,
			// it's a huge number too large to fit in an int. Reject it.
			if !strings.ContainsAny(text, ".eE") {
				return nil, fmt.Errorf("integer overflow: %q", text)
			}
			n.IsFloat = true
			n.Float64 = f
			// If a floating-point extraction succeeded, extract the int if needed.
			if !n.IsInt && float64(int64(f)) == f {
				n.IsInt = true
				n.Int64 = int64(f)
			}
			if !n.IsUint && float64(uint64(f)) == f {
				n.IsUint = true
				n.Uint64 = uint64(f)
			}
		}
	}
	if !n.IsInt && !n.IsUint && !n.IsFloat {
		return nil, fmt.Errorf("illegal number syntax: %q", text)
	}
	return n, nil
}

// NewIdentifier returns a new IdentifierNode with the given identifier name.
func (t *Template) newIdentifier(ident string, pos Pos, line int) *IdentifierNode {
	return &IdentifierNode{NodeBase: NodeBase{TemplateName: t.Name, NodeType: NodeIdentifier, Pos: pos, Line: line}, Ident: ident}
}