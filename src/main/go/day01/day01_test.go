package day01

import "testing"

const rawInput = "180\n152\n159\n171\n178\n169\n212\n213\n214\n222\n228\n215\n228\n240\n248\n220\n224\n201\n212\n218\n217\n225\n218\n255\n256\n260\n261\n262\n263\n254\n255\n261\n270\n248\n252\n258\n259\n243\n242\n240\n233\n241\n250\n256\n258\n256\n258\n261\n263\n274\n262\n248\n265\n266\n276\n279\n273\n274\n273\n275\n280\n281\n282\n284\n285\n286\n289\n294\n310\n313\n318\n327\n342\n346\n358\n378\n386\n391\n377\n378\n381\n378\n377\n369\n372\n375\n368\n377\n354\n319\n314\n337\n338\n351\n356\n361\n345\n346\n337\n338\n339\n365\n370\n368\n383\n384\n399\n416\n417\n422\n427\n431\n434\n438\n439\n440\n438\n437\n439\n444\n446\n449\n461\n489\n483\n466\n471\n472\n474\n478\n476\n481\n482\n483\n485\n510\n512\n519\n525\n539\n562\n576\n583\n587\n607\n605\n601\n616\n615\n630\n635\n638\n643\n644\n642\n652\n662\n661\n687\n683\n682\n701\n710\n712\n721\n722\n719\n726\n734\n729\n728\n752\n754\n758\n765\n764\n765\n766\n781\n790\n798\n799\n801\n802\n801\n802\n830\n831\n832\n844\n858\n850\n847\n846\n853\n854\n853\n859\n855\n856\n858\n848\n880\n859\n866\n867\n865\n868\n878\n877\n878\n883\n894\n900\n909\n912\n896\n897\n885\n886\n885\n893\n901\n903\n900\n896\n904\n912\n914\n905\n904\n912\n925\n924\n935\n949\n950\n947\n949\n950\n946\n948\n947\n946\n944\n941\n974\n982\n981\n1005\n1006\n1003\n1019\n1020\n1019\n1033\n1034\n1035\n1046\n1049\n1043\n1054\n1060\n1059\n1062\n1065\n1066\n1071\n1078\n1080\n1081\n1085\n1090\n1091\n1094\n1099\n1117\n1119\n1120\n1126\n1128\n1126\n1143\n1145\n1140\n1144\n1137\n1110\n1121\n1119\n1120\n1121\n1125\n1116\n1118\n1116\n1105\n1106\n1107\n1105\n1098\n1099\n1124\n1129\n1139\n1143\n1144\n1145\n1148\n1149\n1155\n1174\n1175\n1178\n1174\n1182\n1183\n1190\n1192\n1190\n1191\n1192\n1201\n1202\n1197\n1200\n1201\n1230\n1232\n1250\n1251\n1257\n1274\n1290\n1320\n1319\n1326\n1328\n1349\n1341\n1345\n1355\n1353\n1356\n1359\n1368\n1370\n1387\n1385\n1402\n1413\n1416\n1417\n1422\n1421\n1420\n1413\n1420\n1421\n1422\n1416\n1425\n1431\n1433\n1434\n1436\n1434\n1440\n1458\n1459\n1464\n1466\n1475\n1477\n1500\n1502\n1504\n1492\n1491\n1495\n1497\n1504\n1519\n1516\n1523\n1527\n1528\n1538\n1543\n1544\n1548\n1549\n1556\n1557\n1566\n1564\n1565\n1566\n1572\n1583\n1580\n1579\n1599\n1603\n1615\n1616\n1615\n1618\n1622\n1619\n1630\n1629\n1630\n1629\n1632\n1635\n1636\n1638\n1635\n1634\n1646\n1649\n1651\n1649\n1650\n1648\n1659\n1686\n1696\n1694\n1699\n1700\n1709\n1703\n1702\n1703\n1712\n1716\n1717\n1718\n1705\n1706\n1709\n1713\n1725\n1724\n1750\n1751\n1754\n1752\n1745\n1744\n1750\n1751\n1763\n1764\n1762\n1763\n1769\n1747\n1746\n1737\n1735\n1736\n1735\n1732\n1733\n1736\n1737\n1742\n1748\n1754\n1755\n1736\n1742\n1735\n1737\n1755\n1764\n1762\n1778\n1793\n1815\n1826\n1830\n1831\n1836\n1860\n1858\n1831\n1832\n1833\n1832\n1836\n1850\n1851\n1863\n1865\n1867\n1868\n1875\n1895\n1894\n1895\n1901\n1906\n1938\n1967\n1960\n1961\n1963\n1952\n1958\n1960\n1962\n1960\n1962\n1955\n1963\n1957\n1956\n1959\n1960\n1954\n1955\n1954\n1986\n1985\n2013\n2011\n2012\n2034\n2038\n2041\n2043\n2044\n2045\n2050\n2051\n2061\n2059\n2057\n2064\n2083\n2079\n2084\n2075\n2079\n2082\n2119\n2139\n2147\n2151\n2152\n2148\n2149\n2152\n2155\n2161\n2162\n2185\n2196\n2200\n2202\n2213\n2214\n2215\n2207\n2204\n2208\n2209\n2211\n2205\n2206\n2235\n2255\n2261\n2274\n2275\n2280\n2281\n2288\n2299\n2300\n2301\n2302\n2290\n2309\n2323\n2324\n2329\n2331\n2330\n2317\n2318\n2315\n2302\n2305\n2311\n2316\n2321\n2323\n2332\n2331\n2350\n2342\n2351\n2352\n2356\n2362\n2393\n2400\n2401\n2405\n2421\n2423\n2445\n2441\n2449\n2451\n2454\n2458\n2466\n2463\n2471\n2477\n2478\n2474\n2473\n2475\n2473\n2471\n2472\n2483\n2486\n2489\n2502\n2503\n2504\n2516\n2520\n2523\n2527\n2526\n2530\n2531\n2532\n2536\n2528\n2535\n2542\n2553\n2556\n2560\n2564\n2549\n2542\n2544\n2580\n2581\n2576\n2602\n2601\n2605\n2629\n2628\n2629\n2628\n2629\n2631\n2632\n2624\n2623\n2632\n2641\n2642\n2651\n2661\n2676\n2706\n2716\n2729\n2736\n2737\n2733\n2734\n2740\n2731\n2739\n2734\n2733\n2732\n2733\n2735\n2706\n2709\n2715\n2713\n2716\n2718\n2726\n2728\n2742\n2730\n2731\n2733\n2736\n2761\n2764\n2771\n2772\n2773\n2781\n2782\n2785\n2791\n2793\n2796\n2797\n2798\n2805\n2806\n2811\n2820\n2818\n2816\n2817\n2818\n2806\n2815\n2819\n2824\n2829\n2833\n2834\n2825\n2835\n2837\n2833\n2831\n2833\n2827\n2833\n2848\n2846\n2843\n2844\n2845\n2847\n2857\n2855\n2875\n2877\n2879\n2882\n2891\n2894\n2897\n2899\n2901\n2912\n2917\n2927\n2928\n2921\n2922\n2930\n2931\n2932\n2933\n2949\n2951\n2947\n2953\n2930\n2913\n2930\n2931\n2930\n2932\n2939\n2937\n2939\n2950\n2951\n2953\n2964\n2969\n2979\n2981\n2980\n2983\n2986\n2993\n2994\n2982\n2985\n2986\n2984\n2964\n2987\n2986\n2988\n2997\n2998\n3005\n3017\n3034\n3037\n3038\n3042\n3045\n3056\n3057\n3063\n3069\n3060\n3059\n3060\n3064\n3066\n3056\n3060\n3062\n3077\n3078\n3079\n3067\n3051\n3052\n3051\n3050\n3055\n3058\n3060\n3079\n3098\n3099\n3089\n3113\n3116\n3115\n3126\n3110\n3093\n3105\n3116\n3123\n3126\n3135\n3137\n3139\n3157\n3159\n3192\n3193\n3194\n3192\n3202\n3207\n3215\n3219\n3221\n3222\n3223\n3228\n3222\n3223\n3246\n3256\n3262\n3248\n3252\n3254\n3255\n3256\n3257\n3259\n3260\n3266\n3262\n3251\n3252\n3262\n3268\n3267\n3270\n3266\n3275\n3274\n3256\n3257\n3259\n3262\n3265\n3266\n3270\n3273\n3239\n3252\n3263\n3291\n3294\n3310\n3317\n3359\n3360\n3359\n3360\n3361\n3366\n3369\n3374\n3378\n3376\n3389\n3391\n3405\n3418\n3448\n3455\n3461\n3476\n3485\n3488\n3491\n3492\n3495\n3496\n3497\n3502\n3504\n3507\n3519\n3544\n3546\n3547\n3548\n3549\n3550\n3547\n3556\n3563\n3572\n3577\n3578\n3585\n3593\n3597\n3598\n3602\n3597\n3611\n3590\n3602\n3591\n3592\n3585\n3586\n3578\n3588\n3589\n3590\n3596\n3594\n3606\n3602\n3618\n3619\n3613\n3619\n3597\n3596\n3605\n3606\n3607\n3610\n3612\n3625\n3655\n3657\n3652\n3654\n3660\n3662\n3664\n3684\n3685\n3664\n3667\n3673\n3674\n3697\n3693\n3694\n3697\n3692\n3694\n3697\n3695\n3717\n3731\n3730\n3728\n3730\n3735\n3740\n3738\n3767\n3775\n3776\n3777\n3779\n3783\n3789\n3791\n3793\n3794\n3799\n3800\n3801\n3792\n3798\n3800\n3804\n3778\n3777\n3782\n3790\n3792\n3793\n3816\n3812\n3853\n3860\n3881\n3886\n3888\n3889\n3896\n3899\n3907\n3900\n3905\n3922\n3924\n3940\n3941\n3942\n3930\n3927\n3928\n3933\n3944\n3943\n3956\n3966\n3968\n3972\n3979\n3980\n3982\n3983\n3989\n4010\n4007\n4008\n4019\n4022\n4020\n4011\n4013\n4018\n4029\n4033\n4037\n4038\n4024\n4026\n4014\n4015\n4047\n4064\n4065\n4063\n4065\n4072\n4081\n4082\n4093\n4094\n4099\n4102\n4101\n4100\n4101\n4111\n4115\n4116\n4125\n4139\n4144\n4145\n4129\n4127\n4128\n4129\n4130\n4133\n4148\n4157\n4173\n4174\n4171\n4174\n4175\n4181\n4194\n4199\n4201\n4202\n4207\n4208\n4209\n4210\n4222\n4245\n4243\n4242\n4252\n4255\n4258\n4271\n4288\n4294\n4295\n4297\n4291\n4292\n4304\n4301\n4307\n4310\n4309\n4322\n4324\n4325\n4334\n4337\n4332\n4342\n4345\n4359\n4356\n4378\n4380\n4381\n4387\n4401\n4405\n4409\n4408\n4430\n4419\n4423\n4429\n4434\n4442\n4469\n4476\n4478\n4519\n4520\n4521\n4508\n4514\n4510\n4513\n4515\n4516\n4518\n4524\n4523\n4514\n4518\n4520\n4534\n4537\n4538\n4545\n4546\n4543\n4551\n4566\n4569\n4570\n4586\n4582\n4600\n4599\n4592\n4617\n4621\n4648\n4650\n4654\n4646\n4647\n4660\n4663\n4668\n4679\n4684\n4693\n4695\n4696\n4702\n4701\n4684\n4708\n4694\n4695\n4696\n4723\n4734\n4733\n4736\n4732\n4740\n4736\n4742\n4744\n4754\n4756\n4751\n4755\n4757\n4743\n4738\n4739\n4758\n4763\n4772\n4774\n4775\n4782\n4799\n4808\n4825\n4836\n4839\n4840\n4843\n4838\n4839\n4858\n4865\n4857\n4858\n4855\n4858\n4884\n4889\n4895\n4894\n4902\n4931\n4933\n4934\n4935\n4937\n4938\n4953\n4956\n4959\n4948\n4951\n4960\n4954\n4967\n4971\n4967\n4974\n4975\n4967\n4966\n4962\n4963\n4964\n4967\n4969\n4968\n4974\n4994\n4995\n4981\n4986\n4991\n4998\n4999\n5004\n5011\n5002\n4997\n4989\n4990\n4997\n4993\n5008\n5019\n5024\n5027\n5029\n5038\n5039\n5022\n5013\n5014\n5000\n5002\n5003\n5004\n5003\n5011\n5020\n5023\n5028\n5037\n5063\n5059\n5078\n5079\n5080\n5081\n5078\n5081\n5082\n5087\n5088\n5089\n5107\n5125\n5131\n5142\n5129\n5124\n5127\n5130\n5146\n5149\n5154\n5155\n5156\n5155\n5175\n5174\n5186\n5192\n5201\n5209\n5208\n5213\n5214\n5213\n5216\n5218\n5211\n5202\n5220\n5226\n5225\n5233\n5220\n5221\n5218\n5215\n5220\n5223\n5227\n5231\n5228\n5230\n5231\n5232\n5236\n5215\n5222\n5224\n5241\n5252\n5262\n5273\n5274\n5282\n5283\n5279\n5283\n5289\n5290\n5294\n5296\n5297\n5316\n5314\n5319\n5301\n5302\n5279\n5280\n5278\n5292\n5293\n5294\n5298\n5303\n5316\n5317\n5318\n5324\n5333\n5335\n5336\n5349\n5352\n5355\n5375\n5377\n5378\n5367\n5365\n5366\n5387\n5379\n5385\n5392\n5394\n5395\n5382\n5387\n5386\n5390\n5385\n5400\n5409\n5392\n5390\n5377\n5363\n5352\n5353\n5359\n5376\n5373\n5370\n5373\n5377\n5382\n5395\n5405\n5438\n5449\n5464\n5465\n5468\n5466\n5469\n5454\n5462\n5477\n5475\n5477\n5478\n5465\n5466\n5487\n5464\n5474\n5479\n5498\n5499\n5503\n5518\n5520\n5523\n5534\n5550\n5552\n5576\n5580\n5571\n5574\n5587\n5582\n5565\n5568\n5569\n5576\n5577\n5582\n5595\n5594\n5604\n5587\n5586\n5587\n5591\n5594\n5596\n5600\n5604\n5606\n5614\n5615\n5614\n5616\n5620\n5612\n5630\n5629\n5636\n5639\n5640\n5641\n5611\n5610\n5613\n5596\n5597\n5598\n5599\n5601\n5621\n5628\n5629\n5631\n5634\n5647\n5646\n5623\n5632\n5638\n5649\n5646\n5648\n5670\n5701\n5708\n5724\n5723\n5728\n5762\n5779\n5774\n5777\n5797\n5806\n5817\n5814\n5819\n5844\n5845\n5844\n5847\n5849\n5850\n5860\n5861\n5858\n5860\n5874\n5899\n5900\n5902\n5903\n5904\n5908\n5926\n5928\n5929\n5930\n5918\n5924\n5925\n5927\n5929\n5908\n5912\n5925\n5930\n5927\n5928\n5926\n5937\n5938\n5940\n5939\n5951\n5954\n5957\n5960\n5961\n5969\n5970\n5971\n5972\n5983\n5993\n6009\n6019\n6012\n5993\n6015\n6006\n6020\n6021\n6009\n6020\n6021\n6031\n6033\n6036\n6039\n6038\n6062\n6068\n6069\n6068\n6075\n6093\n6096\n6100\n6116\n6127\n6134\n6136\n6138\n6154\n6155\n6156\n6154\n6150\n6151\n6145\n6151\n6152\n6147\n6166\n6164\n6174\n6185\n6186\n6195\n6172\n6194\n6217\n6226\n6224\n6245\n6251\n6252\n6262\n6266\n6270\n6271\n6276\n6282\n6300\n6301\n6302\n6300\n6319\n6328\n6329\n6330\n6341\n6346\n6351\n6352\n6366\n6365\n6371\n6372\n6380\n6377\n6389\n6399\n6379\n6380\n6382\n6381\n6380\n6384\n6385\n6399\n6400\n6396\n6397\n6405\n6415\n6402\n6403\n6414\n6406\n6399\n6390\n6407\n6409\n6390\n6384\n6386\n6379\n6381\n6388\n6401\n6404\n6397\n6421\n6431\n6434\n6438\n6447\n6446\n6447\n6448\n6457\n6463\n6473\n6476\n6483\n6488\n6486\n6483\n6487\n6491\n6488\n6455\n6471\n6475\n6478\n6498\n6499\n6500\n6501\n6494\n6486\n6487\n6484\n6501\n6497\n6499\n6502\n6504\n6505\n6509\n6525\n6524\n6523\n6543\n6547\n6546\n6536\n6552\n6554\n6562\n6560\n6589\n6591\n6609\n6608\n6604\n6605\n6594\n6608\n6611\n6610\n6611\n6621\n6607\n6612\n6613\n6623\n6628\n6631\n6642\n6663\n6654\n6647\n6640\n6636\n6635\n6641\n6646\n6645\n6646\n6647\n6654\n6655\n6656\n6668\n6669\n6679\n6675\n6687\n6685\n6687\n6716\n6717\n6722\n6730\n6729\n6732\n6735\n6736\n6748\n6759\n6758\n6757\n6780\n6781\n6784\n6794\n6796\n6811\n6817\n6818\n6819\n6820\n6821\n6823\n6812\n6811\n6837\n6839\n6842\n6844\n6842\n6841\n6852\n6853\n6856\n6855\n6865\n6875\n6876\n6898\n6904\n6906\n6910\n6914\n6915\n6905\n6890\n6896\n6894\n6891\n6899\n6904\n6909\n6927\n6928\n6938\n6935\n6936\n6937\n6946\n6949\n6963\n6966\n6975\n6976\n6977\n6973\n6974\n6980\n6984\n7014\n6993\n7004\n7005\n7013\n7016\n7014\n7018\n7029\n7030\n7031\n7038\n7039\n7056\n7079\n7081\n7085\n7103\n7105\n7098\n7101\n7102\n7106\n7111\n7112\n7114\n7101\n7107\n7117\n7085\n7093\n7091\n7119\n7115\n7126\n7130\n7129\n7139\n7142\n7147\n7169\n7160\n7171\n7187\n7188\n7189\n7188\n7190\n7194\n7197\n7212\n7220\n7228\n7227\n7223\n7225\n7236\n7237\n7241\n7243\n7246\n7252\n7280\n7279\n7294\n7295\n7298\n7300\n7289\n7291\n7290\n7288\n7289\n7291\n7290\n7295\n7294\n7275\n7278\n7290\n7295\n7302\n7304\n7309\n7308\n7321\n7325\n7328\n7346\n7338\n7340\n7354\n7361\n7368\n7391\n7415\n7426\n7428\n7429\n7422\n7423\n7425\n7426\n7429\n7431\n7442\n7440\n7443\n7457\n7458\n7457\n7464\n7474\n7477\n7469\n7467\n7471\n7479\n7480\n7485\n7482\n7483\n7484\n7488\n7497"

func TestDay01Question01SampleData(t *testing.T) {
	input := "199\n200\n208\n210\n200\n207\n240\n269\n260\n263\n"
	if day01Question01(input) != 7 {
		t.Fail()
	}
}

func TestDay01Question01ActualData(t *testing.T) {
	if day01Question01(rawInput) != 1529 {
		t.Fail()
	}
}

func TestDay01Question02SampleData(t *testing.T) {
	input := "607\n618\n618\n617\n\n647\n716\n769\n792"
	if day01Question02(input) != 5 {
		t.Fail()
	}
}

func TestDay01Question02ActualData(t *testing.T) {
	if day01Question02(rawInput) != 1567 {
		t.Fail()
	}
}