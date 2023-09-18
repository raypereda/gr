# gr
### This a rudimentary grep-like command that works on paragraphs.

## Usage
```
Usage of gr:
    gr [-r <string>] <pattern> [<file>]
    gr [-h]
If no files names, it reads from standard input.
Print records that match golang regex pattern.
  -r string
        Record separator (default "(\\n\\n)|(\\r\\n\\r\\n)")
```

The Golang regular expression syntax specification is [here](https://golang.org/pkg/regexp/syntax/)
.

## Runtime

The runtime is measured on a Mac 2021 Macbook with Apple M1 chip.

```
time gr law The_King_James_Version_of_the_Bible.txt > /dev/null  
0.17s user 0.01s system 100% cpu 0.180 total
```

For comparison with mac /usr/bin/grep, here is the timing.:
```
time grep law  The_King_James_Version_of_the_Bible.txt > /dev/null
0.05s user 0.00s system 96% cpu 0.061 total
```


## Installation 

These are 64-bit statically linked executables:
1. gr-linux-intel
2. gr-mac-apple-silicon
3. gr-mac-intel
4. gr.exe-windows

To install, copy to a directory that on your shell $PATH environment variable.

## Example 1
```
$ gr critic Man_in_the_Arena--Teddy_Roosevelt.txt
```
```
Man_in_the_Arena--Teddy_Roosevelt.txt
record separator:  (\n\n)|(\r\n\r\n)
37  paragraphs
Match #  1
Let the man of learning, the man of lettered leisure, beware of that queer and 
cheap temptation to pose to himself and to others as a cynic, as the man who 
has outgrown emotions and beliefs, the man to whom good and evil are as one. 
The poorest way to face life is to face it with a sneer. There are many men who 
feel a kind of twisted pride in cynicism; there are many who confine themselves 
to criticism of the way others do what they themselves dare not even attempt. 
There is no more unhealthy being, no man less worthy of respect, than he who 
either really holds, or feigns to hold, an attitude of sneering disbelief 
toward all that is great and lofty, whether in achievement or in that noble 
effort which, even if it fails, comes second to achievement. A cynical habit of 
thought and speech, a readiness to criticize work which the critic himself 
never tries to perform, an intellectual aloofness which will not accept contact 
with life’s realities—all these are marks, not as the possessor would fain to 
think, of superiority, but of weakness. They mark the men unfit to bear their 
part painfully in the stern strife of living, who seek, in the affectation of 
contempt for the achievement of others, to hide from others and from themselves 
their own weakness. The role is easy; there is none easier, save only the role 
of the man who sneers alike at both criticism and performance.
Match #  2
It is not the critic who counts; not the man who points out how the strong man 
stumbles, or where the doer of deeds could have done them better. The credit 
belongs to the man who is actually in the arena, whose face is marred by dust 
and sweat and blood; who strives valiantly; who errs, who comes short again and 
again, because there is no effort without error and shortcoming; but who does 
actually strive to do the deeds; who knows the great enthusiasms, the great 
devotions; who spends himself in a worthy cause; who at the best knows in the 
end the triumph of high achievement, and who at the worst, if he fails, at 
least fails while daring greatly, so that his place shall never be with those 
cold and timid souls who neither know victory nor defeat. Shame on the man of 
cultivated taste who permits refinement to develop into fastidiousness that 
unfits him for doing the rough work of a workaday world. Among the free peoples 
who govern themselves there is but a small field of usefulness open for the men 
of cloistered life who shrink from contact with their fellows. Still less room 
is there for those who deride or slight what is done by those who actually bear 
the brunt of the day; nor yet for those others who always profess that they 
would like to take action, if only the conditions of life were not exactly what 
they actually are. The man who does nothing cuts the same sordid figure in the 
pages of history, whether he be cynic, or fop, or voluptuary. There is little 
use for the being whose tepid soul knows nothing of the great and generous 
emotion, of the high pride, the stern belief, the lofty enthusiasm, of the men 
who quell the storm and ride the thunder. Well for these men if they succeed; 
well also, though not so well, if they fail, given only that they have nobly 
ventured, and have put forth all their heart and strength. It is war-worn 
Hotspur, spent with hard fighting, he of the many errors and the valiant end, 
over whose memory we love to linger, not over the memory of the young lord who 
“but for the vile guns would have been a valiant soldier.”
2  matches
```

## Example 2
```
$ gr garden Candide_by_Voltaire.txt 
```

```
Candide_by_Voltaire.txt
record separator:  (\n\n)|(\r\n\r\n)
967  paragraphs
Match #  1
She took him by the arm, and walked with him about a quarter of a mile
into the country; they arrived at a lonely house, surrounded with
gardens and canals. The old woman knocked at a little door, it opened,
she led Candide up a private staircase into a small apartment richly
furnished. She left him on a brocaded sofa, shut the door and went away.
Candide thought himself in a dream; indeed, that he had been dreaming
unluckily all his life, and that the present moment was the only
agreeable part of it all.
Match #  2
"As soon as my companions could walk, they were obliged to set out for
Moscow. I fell to the share of a Boyard who made me his gardener, and
gave me twenty lashes a day. But this nobleman having in two years' time
been broke upon the wheel along with thirty more Boyards for some broils
at court, I profited by that event; I fled. I traversed all Russia; I
was a long time an inn-holder's servant at Riga, the same at Rostock, at
Vismar, at Leipzig, at Cassel, at Utrecht, at Leyden, at the Hague, at
Rotterdam. I waxed old in misery and disgrace, having only one-half of
my posteriors, and always remembering I was a Pope's daughter. A hundred
times I was upon the point of killing myself; but still I loved life.
This ridiculous foible is perhaps one of our most fatal characteristics;
for is there anything more absurd than to wish to carry continually a
burden which one can always throw down? to detest existence and yet to
cling to one's existence? in brief, to caress the serpent which devours
us, till he has eaten our very heart?
Match #  3

Candide and Martin went in a gondola on the Brenta, and arrived at the
palace of the noble Signor Pococurante. The gardens, laid out with
taste, were adorned with fine marble statues. The palace was beautifully
built. The master of the house was a man of sixty, and very rich. He
received the two travellers with polite indifference, which put Candide
a little out of countenance, but was not at all disagreeable to Martin.
Match #  4
After their survey of the library they went down into the garden, where
Candide praised its several beauties.
Match #  5
It is natural to imagine that after so many disasters Candide married,
and living with the philosopher Pangloss, the philosopher Martin, the
prudent Cacambo, and the old woman, having besides brought so many
diamonds from the country of the ancient Incas, must have led a very
happy life. But he was so much imposed upon by the Jews that he had
nothing left except his small farm; his wife became uglier every day,
more peevish and unsupportable; the old woman was infirm and even more
fretful than Cunegonde. Cacambo, who worked in the garden, and took
vegetables for sale to Constantinople, was fatigued with hard work, and
cursed his destiny. Pangloss was in despair at not shining in some
German university. For Martin, he was firmly persuaded that he would be
as badly off elsewhere, and therefore bore things patiently. Candide,
Martin, and Pangloss sometimes disputed about morals and metaphysics.
They often saw passing under the windows of their farm boats full of
Effendis, Pashas, and Cadis, who were going into banishment to Lemnos,
Mitylene, or Erzeroum. And they saw other Cadis, Pashas, and Effendis
coming to supply the place of the exiles, and afterwards exiled in their
turn. They saw heads decently impaled for presentation to the Sublime
Porte. Such spectacles as these increased the number of their
dissertations; and when they did not dispute time hung so heavily upon
their hands, that one day the old woman ventured to say to them:
Match #  6
"I do not know," answered the worthy man, "and I have not known the name
of any Mufti, nor of any Vizier. I am entirely ignorant of the event you
mention; I presume in general that they who meddle with the
administration of public affairs die sometimes miserably, and that they
deserve it; but I never trouble my head about what is transacting at
Constantinople; I content myself with sending there for sale the fruits
of the garden which I cultivate."
Match #  7
"I know also," said Candide, "that we must cultivate our garden."
Match #  8
"All that is very well," answered Candide, "but let us cultivate our
garden."
8  matches
```

