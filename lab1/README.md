### How to run this code on Linux?
> swipl \
consult('main.pl').

or 
https://swish.swi-prolog.org

### Information about facts and rules
- **21** facts with single param
- **17** facts with 2 params
- **5** rules

### Query examples
- child('Mask of Madness', 'Quarterstaff'). -  ***true***
- can_be_assembled('Mask of Madness', 'Quarterstaff'). -  ***false***
- can_be_assembled('Sages Mask', 'Vladmirs Offering'). -  ***true***
- contains_element('Helm of the Overlord', 'Ring of Protection'). -  ***true***
- can_be_dissasembled_into('Mask of Madness', 'Quarterstaff'). -  ***true***
- recipe_for_item('Helm of the Dominator Recipe', 'Helm of the Dominator'). - ***true***