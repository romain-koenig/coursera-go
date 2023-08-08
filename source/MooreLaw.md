## Moore's Law

**[Moore's Law](https://en.wikipedia.org/wiki/Moore%27s_law)** refers to an observation made by [Gordon Moore](https://en.wikipedia.org/wiki/Gordon_Moore), co-founder of Intel, in 1965. Contrary to its name, it's not a physical law but an empirical observation. Moore observed that the transistor density on a microchip was doubling approximately every two years, leading to an exponential increase in computing power. This prediction held true for several decades.

## Why Moore's Law Has Stopped Being True

Several factors, both physical and economic in nature, have constrained the continued applicability of Moore's Law:

- **Thermal Constraints**: As we pack more transistors onto a chip, dissipating the generated heat becomes a significant challenge. Excessive heat can lead to reduced performance and, in extreme cases, potential damage (or even melting) of the chip.

- **Power Consumption**: While smaller transistors have been historically associated with lower power consumption, there's a limit to how low this can go. The equation P = αCFV² shows that dynamic power is proportional to the square of the voltage swing. This makes Voltage the number one factor in determining power consumption.
Unfortunately, this voltage can't be reduced indefinitely because:
  - It needs to be sufficient to ensure the transistor switches effectively.
  - A lower threshold makes it difficult to differentiate between values in the presence of noise (and there is always noise in a real world scenario).
  
  Additionally, as components get smaller, leakage increases due to diminishing isolation between them.

- **Other Constraints**: Some limitations, though not extensively covered in the courses, are critical. These include:
  - **Industrial Constraints**: Semiconductor companies face challenges in miniaturizing components further.
  - **Physical Limitations**: There's an inherent limit to how small components can be manufactured while remaining functional.

## New Directions

As the pace of Moore's Law slows, various alternatives are being pursued to continue improving computing power:

- **Multi-core processors**: Instead of solely increasing transistor density, the focus has shifted to increasing the number of cores per chip.

- **Specialized Hardware**: There's an increasing trend towards specialized hardware designed for specific tasks. Cryptocurrencies mining and AI are two areas where this is particularly evident.
